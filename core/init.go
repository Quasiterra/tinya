package core

import "github.com/aws/aws-sdk-go/aws/session"

var (
	awsSession = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
)
