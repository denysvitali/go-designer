package designer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
)

type Client struct {
	token     string
	sessionId string
}

func New(token string) Client {
	return Client{token: token}
}

type CustomRoundTripper struct {
}

func (c CustomRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return http.DefaultTransport.RoundTrip(request)
}

var _ http.RoundTripper = (*CustomRoundTripper)(nil)

func (c *Client) GenerateImages(prompt string) (*Response, error) {
	// Form data parameters
	formData := map[string]string{
		"dalle-caption":               prompt,
		"dalle-scenario-name":         "TextToImage",
		"dalle-batch-size":            "3",
		"dalle-image-response-format": "UrlWithBase64Thumbnail",
		"dalle-seed":                  "17",
	}

	// Create a new multipart writer
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Add form data parameters to the request
	for key, value := range formData {
		header := make(textproto.MIMEHeader)
		header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`, key))
		part, err := writer.CreatePart(header)
		if err != nil {
			fmt.Println("Error creating multipart form part:", err)
			return nil, nil
		}
		part.Write([]byte(value))
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	// Send the HTTP POST request
	url := "https://designerapp.officeapps.live.com/designerapp/DallE.ashx?action=GetDallEImagesCogSci"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("SessionId", uuid.New().String())

	err = c.addAuthentication(req)
	if err != nil {
		return nil, err
	}

	httpClient := http.Client{Transport: CustomRoundTripper{}}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, nil
	}

	// Process the JSON response
	var jsonResponse Response
	err = json.Unmarshal(responseBody, &jsonResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return nil, nil
	}
	return &jsonResponse, nil
}

func SaveImages(jsonResponse *Response, baseName string) ([]string, error) {

	// Access the image URLs and thumbnail data
	var fileNames []string
	thumbnailData := jsonResponse.ImageUrlsThumbnail
	for i, entry := range thumbnailData {
		fileName := fmt.Sprintf(
			"%s-%d.jpg",
			baseName,
			i,
		)
		err := saveImage(
			entry.ImageUrl,
			path.Join(
				"output",
				fileName,
			),
		)
		if err != nil {
			return nil, err
		}
		fileNames = append(fileNames, fileName)
	}
	return fileNames, nil
}

func saveImage(imageUrl string, fileName string) error {
	res, err := http.Get(imageUrl)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = io.Copy(f, res.Body)
	return err
}

func (c *Client) addAuthentication(req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	return nil
}
