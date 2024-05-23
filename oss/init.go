package oss

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func init() {
	endpoint := "139.9.128.19:9000"
	accessKeyID := "A0BO5jFTMMekd8Jvf0pQ"
	secretAccessKey := "LQnw46gMgZHvEnFLmBpxTakS6mvXjbJLJlXaH0a5"
	bucketName := "zhulong"

	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		panic(err)
	}

	minioClient = client

	if err := CreateBucket(bucketName); err != nil {
		panic(err)
	}
}
