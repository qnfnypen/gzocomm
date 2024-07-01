package aws3

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewService 创建 aws s3 服务
func NewService(conf *Config) *Service {
	creds := credentials.NewStaticCredentials(conf.AccessKeyID, conf.SecretAccessKey, conf.SessionToken)
	session := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: creds,
	}))

	return &Service{
		s3Svc:  s3.New(session),
		bucket: conf.Bucket,
	}
}

// UploadLocalFile 上传本地文件，dir 指定s3上存储的目录，返回的为s3上存储的key，需要加上endpoint或者cdn前缀才可以访问
func (svc *Service) UploadLocalFile(fn string, dir ...string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", fmt.Errorf("open loacl file error:%w", err)
	}
	defer f.Close()

	ext := filepath.Ext(f.Name())

	return svc.upload(f, "", ext, dir...)
}

// UploadFormFile 上传form文件，dir 指定s3上存储的目录，返回的为s3上存储的key，需要加上endpoint或者cdn前缀才可以访问
func (svc *Service) UploadFormFile(fh *multipart.FileHeader, dir ...string) (string, error) {
	f, err := fh.Open()
	if err != nil {
		return "", fmt.Errorf("open form file: %w", err)
	}
	defer f.Close()

	ctype := fh.Header.Get("content-type")
	ext := filepath.Ext(fh.Filename)

	return svc.upload(f, ctype, ext, dir...)
}

// DeleteFile 从s3上删除文件
func (svc *Service) DeleteFile(key string) error {
	_, err := svc.s3Svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(svc.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("delete file from s3 error:%w", err)
	}

	return nil
}
