package helpers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sts"
)

type Credentials struct {
	AccessKey string
	SecretKey string
	Region    string
}

type SqsClient struct {
	sqsClient *sqs.SQS
}

func NewSqsClient(creds *Credentials) (*SqsClient, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(creds.Region)}, nil)
	if err!= nil {
		return nil, err
	}

	return &SqsClient{
		sqsClient: sqs.New(sess),
	}, nil
}

func (c *SqsClient) SendMessage(queueURL string, message string) error {
	_, err := c.sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:   aws.String(queueURL),
	})
	return err
}

func GetToken(creds *Credentials) (string, error) {
	stsClient := sts.New(session.New(&aws.Config{Region: aws.String(creds.Region)}, nil))
	resp, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err!= nil {
		return "", err
	}

	return *resp.Arn, nil
}

func GetRequestID(r *http.Request) string {
	return r.Header.Get("X-Request-Id")
}

func GetContextTimeout(ctx context.Context) time.Duration {
	return ctx.Deadline().Sub(time.Now())
}