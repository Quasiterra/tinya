package core

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Svc = s3.New(awsSession)

func ListBuckets() (string, error) {
	res, err := s3Svc.ListBuckets(nil)
	return res.GoString(), err
}
