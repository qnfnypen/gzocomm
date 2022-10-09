package rekog

import "github.com/aws/aws-sdk-go/service/rekognition"

// Config 服务配置参数，是aws.Config的精简配置，详情说明可以查看aws.Config
type Config struct {
	Region          string  // required, http://docs.aws.amazon.com/general/latest/gr/rande.html
	Bucket          string  // 如果是s3图片识别，则需要
	MaxLabels       int64   // 如果是图片标签提取，则需要
	MinConfidence   float64 // required, 图片鉴黄指标
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// Service 亚马逊图片识别服务
type Service struct {
	rgSvc         *rekognition.Rekognition
	bucket        string // 文件保存的桶，通常根据项目划分
	MaxLabels     int64
	MinConfidence float64
}
