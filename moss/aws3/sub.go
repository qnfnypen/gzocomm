package aws3

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/qnfnypen/gzocomm/moss/helper"
)

func (svc *Service) upload(f io.ReadSeeker, ctype, ext string, dir ...string) (string, error) {
	// 获取文件 content-type
	ctype, ext, _ = helper.GetContentTypeAndExt(f, ctype, ext)

	// 根据文件内容生成唯一hash作为key
	f.Seek(0, io.SeekStart) // rewind to output whole file
	h := md5.New()
	io.Copy(h, f)
	key := hex.EncodeToString(h.Sum(nil)) + ext
	if len(dir) > 0 {
		key = fmt.Sprintf("%s/%s", strings.TrimSuffix(dir[0], "/"), key)
	}
	f.Seek(0, io.SeekStart) // rewind to output whole file

	// 上传到s3上
	_, err := svc.s3Svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(svc.bucket),
		Key:         aws.String(key),
		Body:        f,
		ContentType: aws.String(ctype),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case request.CanceledErrorCode:
				return "", fmt.Errorf("upload canceled due to timeout:%w", aerr)
			default:
				return "", errors.New(aerr.Error())
			}
		}

		return "", fmt.Errorf("failed to upload object:%w", err)
	}

	return key, nil
}
