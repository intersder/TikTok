package services

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"vibrato/config"
)

var UploadService = newUploadService()

func newUploadService() *uploadService {
	return &uploadService{}
}

type uploadService struct {
}

func (s uploadService) GetSnapshot(finalName string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}
	resp, err := client.CI.GetSnapshot(context.Background(), finalName, &cos.GetSnapshotOptions{
		Time: 1,
	})
	if err != nil {
		return "", err
	}

	response, err := client.Object.Put(context.Background(), finalName+".jpg", resp.Body, nil)
	if err != nil {
		return "", err
	}
	return response.Request.URL.String(), nil
}

func (s uploadService) Upload(file *multipart.FileHeader, finalName string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}

	fd, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fd.Close()

	response, err := client.Object.Put(context.Background(), finalName, fd, nil)
	if err != nil {
		return "", err
	}

	return response.Request.URL.String(), nil
}

func getClient() (*cos.Client, error) {

	CosBucket := config.Config.CosBucket
	CosRegion := config.Config.CosRegion
	u, err := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", CosBucket, CosRegion))
	if err != nil {
		return nil, err
	}
	su, err := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", CosRegion))
	if err != nil {
		return nil, err
	}
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Config.CosSecretId,  // 替换为用户的 SecretId，
			SecretKey: config.Config.CosSecretKey, // 替换为用户的 SecretKey
		},
	})
	return client, nil
}
