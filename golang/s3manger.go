package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func s3manger() {
	bucket := "xipe-electronic-letter-files-env1"
	key := "YOUR_OBJECT_KEY"
	fileName := "test.pdf"

	cf := aws.Credentials{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID_xipe"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY_xipe"),
	}

	cfg, err := external.LoadDefaultAWSConfig(
		external.WithCredentialsValue(cf),
	)

	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := s3.New(cfg)
	ctx := context.Background()
	req := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	resp, err := req.Send(ctx)
	if err != nil {
		panic(err)
	}

	s3objectBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// create file
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	bytesWritten, err := f.Write(s3objectBytes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Fetched %d bytes for S3Object\n", bytesWritten)
	fmt.Printf("successfully downloaded data from %s/%s\n to file %s", bucket, key, fileName)
}
