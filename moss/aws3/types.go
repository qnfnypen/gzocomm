package aws3

import "github.com/aws/aws-sdk-go/service/s3"

// Config 服务配置参数，是aws.Config的精简配置，详情说明可以查看aws.Config
type Config struct {
	Region          string // required, http://docs.aws.amazon.com/general/latest/gr/rande.html
	Bucket          string // required
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// Service 亚马逊 s3 服务
type Service struct {
	s3Svc  *s3.S3
	bucket string // 文件保存的桶，通常根据项目划分
}
