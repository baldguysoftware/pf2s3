package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/urfave/cli"
)

var app *cli.App

const VERSION = "0.8"

func main() {
	app = cli.NewApp()
	app.Name = "pf2s3"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "bucket,b",
			Usage:  "The bucket to upload to",
			EnvVar: "S3_BUCKET",
		},
		cli.StringFlag{
			Name:   "path,p",
			Usage:  "The path in the bucket to upload to",
			EnvVar: "BUCKET_PATH",
			Value:  "pf2s3",
		},
		cli.StringFlag{
			Name:   "region,r",
			Usage:  "The region",
			EnvVar: "AWS_REGION",
			Value:  "us-east-1",
		},
	}
	app.Action = sendToS3
	app.Run(os.Args)
}

func sendToS3(c *cli.Context) error {
	bucket := c.String("bucket")
	base_path := c.String("path")
	region := c.String("region")
	messageB, err := ioutil.ReadAll(os.Stdin)
	if len(messageB) < 32 {
		log.Fatal("Message too small to be an email")
	}
	mreader := bytes.NewReader(messageB)
	m, err := mail.ReadMessage(mreader)
	if err != nil {
		log.Fatal(err)
	}

	tohdr, err := mail.ParseAddress(m.Header.Get("To"))
	if err != nil {
		log.Fatal(err)
	}
	fromhdr, err := mail.ParseAddress(m.Header.Get("From"))
	if err != nil {
		log.Fatal(err)
	}
	tags := fmt.Sprintf("sender=%s&recipient=%s", fromhdr.Address, tohdr.Address)

	s, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		log.Fatal(err)
	}
	mreader.Seek(0, 0)

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(base_path + "/" + m.Header.Get("Message-Id")),
		Body:                 mreader,
		ServerSideEncryption: aws.String("AES256"),
		Tagging:              aws.String(tags),
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
