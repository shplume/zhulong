package oss

import (
	"context"
	"io"
	"mime"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
)

const normalContentType = "application/octet-stream"
const ExpireTime = 3600 // 1h

func CreateBucket(bucketName string) error {
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func UploadFile(bucketName string, objectName string, file io.Reader, size int64) (info minio.UploadInfo, err error) {
	opts := minio.PutObjectOptions{}

	if opts.ContentType = mime.TypeByExtension(filepath.Ext(objectName)); opts.ContentType == "" {
		opts.ContentType = normalContentType
	}

	return minioClient.PutObject(context.Background(), bucketName, objectName, file, size, opts)
}

func GetFileTemporaryURL(bucketName, objectName string) (string, error) {
	expires := time.Second * time.Duration(ExpireTime)

	urlPath, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, expires, nil)

	if err != nil {
		return "", err
	}

	return urlPath.String(), nil
}
