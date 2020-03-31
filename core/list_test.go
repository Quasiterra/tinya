package core

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestListBuckets(t *testing.T) {
	if _, err := ListBuckets(); err != nil {
		t.Error(err)
	}
}

func TestListBucketsError(t *testing.T) {
	tempS3 := s3Svc
	s3Svc = s3.New(session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("not", "a", "thing"),
	})))
	if res, err := ListBuckets(); err == nil {
		fmt.Println(res)
		t.Error("Bad credentials! Should fail!")
	}
	s3Svc = tempS3
}
