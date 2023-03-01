package ss3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func hello() {
	creds := credentials.NewStaticCredentials("0JOR3AT489XN82R9KBFQ", "zzZN7TEI9xmfL8ga0MPDANyLSYtBWTeVriCIytlw", "")
	cfg := aws.NewConfig().
		WithRegion("cn").
		WithEndpoint("http://192.168.50.76:18080").
		WithLogLevel(aws.LogDebugWithHTTPBody | aws.LogDebugWithRequestRetries).
		WithS3ForcePathStyle(true).
		WithDisableSSL(true).WithCredentials(creds)
	sess := session.Must(session.NewSession(cfg))
	svc := s3.New(sess)
	res, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		fmt.Println(res, err)
		return
	}
	fmt.Println(res)
}
