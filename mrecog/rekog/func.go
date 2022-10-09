package rekog

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// NewService 创建 aws rekognition 服务
func NewService(conf *Config) *Service {
	creds := credentials.NewStaticCredentials(conf.AccessKeyID, conf.SecretAccessKey, conf.SessionToken)
	session := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: creds,
	}))

	return &Service{
		rgSvc:         rekognition.New(session),
		bucket:        conf.Bucket,
		MaxLabels:     conf.MaxLabels,
		MinConfidence: conf.MinConfidence,
	}
}

// DetectModerationLabelsByS3 根据亚马逊s3的key，进行图片违规监测，只支持jpeg和png图片
// 返回监测的: 结果，是否正常，错误
func (svc *Service) DetectModerationLabelsByS3(name string) (*rekognition.DetectModerationLabelsOutput, bool, error) {
	input := &rekognition.DetectModerationLabelsInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(svc.bucket),
				Name:   aws.String(name),
			},
		},
		MinConfidence: aws.Float64(svc.MinConfidence),
	}

	return svc.detectModerationLabels(input)
}

// DetectModerationLabelsByBody 上传图片body(最大为5MBs)，进行图片违规监测，只支持jpeg和png图片
// 返回监测的: 结果，是否正常，错误
func (svc *Service) DetectModerationLabelsByBody(body []byte) (*rekognition.DetectModerationLabelsOutput, bool, error) {
	input := &rekognition.DetectModerationLabelsInput{
		Image: &rekognition.Image{
			Bytes: body,
		},
		MinConfidence: aws.Float64(svc.MinConfidence),
	}

	return svc.detectModerationLabels(input)
}
