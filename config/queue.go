package config

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

// Fuunção para inicializar uma sessão na aws
func NewSession() *sqs.SQS {
	LogInfo("init sqs session", zap.String("journey", "initAWSSession"))
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Getenv("SQS_ENDPOINT")),
		Region:   aws.String(os.Getenv("AWS_REGION")),
	}))
	return sqs.New(sess)
}

func ReceiveMessages(svc *sqs.SQS) []*sqs.Message {
	LogInfo("listening the queue to receive messages", zap.String("journey", "receiveMessages"))
	receiveParams := &sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(1),
		QueueUrl:            aws.String(os.Getenv("SQS_QUEUE_URI")),
		WaitTimeSeconds:     aws.Int64(20),
	}

	result, err := svc.ReceiveMessage(receiveParams)

	if err != nil {
		time.Sleep(1 * time.Second)
		LogError("Error to receive messages", err)
		return nil
	}

	return result.Messages
}

func DeleteMessage(svc *sqs.SQS, message *sqs.Message) {
	LogInfo("delete messages", zap.String("journey", "deleteMessages"))
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(os.Getenv("SQS_QUEUE_URI")),
		ReceiptHandle: message.ReceiptHandle,
	}

	_, err := svc.DeleteMessage(deleteParams)

	if err != nil {
		LogError("Error to delete messages", err)
	}
}
