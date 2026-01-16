package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an SQS client
	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String("http://localstack:4566")
	})

	queueName := "my-test-queue"
	queue, err := client.CreateQueue(ctx, &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		log.Fatalf("キュー作成エラー: %v", err)
	}

	queueURL := *queue.QueueUrl
	log.Println("Producer 開始...")

	for i := 1; ; i++ {
		msg := fmt.Sprintf("メッセージ %d from Producer", i)
		_, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
			QueueUrl:    &queueURL,
			MessageBody: &msg,
		})
		if err != nil {
			log.Printf("送信エラー: %v", err)
		}

		log.Printf("送信完了: %s", msg)
		time.Sleep(2 * time.Second)
	}
}
