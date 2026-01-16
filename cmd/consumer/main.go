package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	ctx := context.Background()
	cfg, _ := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))

	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String("http://localstack:4566")
	})

	// ConsumerはURLが固定で分かっている前提（またはCreateを呼んでも良い）
	queueUrl := "http://localstack:4566/000000000000/my-test-queue"

	log.Println("Consumer起動... 受信待機中")
	for {
		output, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueUrl),
			MaxNumberOfMessages: 1,
			WaitTimeSeconds:     10,
		})
		if err != nil {
			log.Printf("受信エラー: %v", err)
			continue
		}

		for _, m := range output.Messages {
			log.Printf("【受信】: %s", *m.Body)
			_, err := client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueUrl),
				ReceiptHandle: m.ReceiptHandle,
			})
			if err != nil {
				log.Printf("削除エラー: %v", err)
			}

			log.Printf("削除完了")
		}
	}
}
