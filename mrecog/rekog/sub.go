package rekog

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func (svc *Service) detectModerationLabels(input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, bool, error) {
	result, err := svc.rgSvc.DetectModerationLabels(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rekognition.ErrCodeInvalidS3ObjectException:
				return nil, false, fmt.Errorf("%s:%w", rekognition.ErrCodeInvalidS3ObjectException, aerr)
			case rekognition.ErrCodeInvalidImageFormatException:
				return nil, false, fmt.Errorf("%s:%w", rekognition.ErrCodeInvalidImageFormatException, aerr)
			case rekognition.ErrCodeThrottlingException:
				return nil, false, fmt.Errorf("%s:%w", rekognition.ErrCodeThrottlingException, aerr)
			default:
				return nil, false, errors.New(aerr.Error())
			}
		} else {
			return nil, false, fmt.Errorf("detect Moderation Labels error:%w", err)
		}
	}

	if len(result.ModerationLabels) == 0 {
		return result, true, nil
	}

	return result, false, nil
}
