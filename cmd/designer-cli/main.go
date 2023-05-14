package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/alexflint/go-arg"
	"github.com/denysvitali/go-designer"
	"github.com/sirupsen/logrus"
)

var args struct {
	Prompt    string `arg:"positional,required"`
	Token     string `arg:"-t,--token,env:DESIGNER_TOKEN,required"`
}

var logger = logrus.New()

func main() {
	doMain()
}

func doMain() {
	arg.MustParse(&args)

	c := designer.New(args.Token)

	images, err := c.GenerateImages(args.Prompt)
	if err != nil {
		logger.Fatalf("unable to generate images: %v", err)
	}

	hasher := md5.New()
	hasher.Write([]byte(args.Prompt))
	promptMd5 := hex.EncodeToString(hasher.Sum(nil))

	fileNames, err := designer.SaveImages(images, promptMd5)
	if err != nil {
		logger.Fatalf("unable to save images: %v", err)
	}

	for _, v := range fileNames {
		logger.Infof("written %s", v)
	}
}
