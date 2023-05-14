# go-designer

A CLI and Go package to interact with 
[Microsoft Designer](https://designer.microsoft.com/) to generate images.

## Example

```bash
designer-cli "Flat vector illustration of a developer using a terminal"
```

Returns (after a few seconds):

```plain
INFO[0016] written 5031222530fccaa508c0a481694ae80f-0.jpg 
INFO[0016] written 5031222530fccaa508c0a481694ae80f-1.jpg 
INFO[0016] written 5031222530fccaa508c0a481694ae80f-2.jpg
```

The result `output/5031222530fccaa508c0a481694ae80f-0.jpg` should return something like this:

![Flat vector illustration of a developer using a terminal](docs/example-1.jpg)

## Installing

### Requirements

- Go 1.20+


### Steps

```bash
go install github.com/denysvitali/go-designer/cmd/designer-cli
designer-cli -h
```

### Getting the token

1. Visit https://designer.microsoft.com/
2. Retrieve the token (two options)
   1. Open the Developer Tools and check the response for the request to:
     `https://login.microsoftonline.com/consumers/oauth2/v2.0/token`
      Copy the value of `access_token` in your clipboard and keep it for later
   2. Generate an image and get the token from the `Authorization: Bearer TOKEN` header
3. Store the token in the `DESIGNER_TOKEN` environment variable as follows:
   ```bash
   read -s -r DESIGNER_TOKEN
   # Paste your token and press ENTER
   export DESIGNER_TOKEN
   ```

### Using the CLI

Pretty straight-forward:

```
designer-cli "Your prompt goes here"
```


### Prompt suggestions

#### Generating Flat Vector Images

```plain
Flat vector illustration of ...
```

Example result:
![Flat vector illustration of a developer using a terminal](docs/example-1.jpg)
