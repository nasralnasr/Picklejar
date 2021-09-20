package utils

import (
  "context"

  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSSubscribeAPI interface {
	Subscribe(ctx context.Context,
		params *sns.SubscribeInput,
		optFns ...func(*sns.Options)) (*sns.SubscribeOutput, error)
}

func subscribeTopic(c context.Context, api SNSSubscribeAPI, input *sns.SubscribeInput) (*sns.SubscribeOutput, error) {
	return api.Subscribe(c, input)
}

func SubscribeToSMS(phoneNumber *string) {
  cfg, err := config.LoadDefaultConfig(context.TODO())
  if err != nil {
    panic("configuration error, " + err.Error())
  }

  client := sns.NewFromConfig(cfg)



  topicARN := aws.String("arn:aws:sns:us-east-1:375351057769:DeliveryNotification")

  input := &sns.SubscribeInput{
    Endpoint:              phoneNumber,
    Protocol:              aws.String("sms"),
    ReturnSubscriptionArn: true, // Return the ARN, even if user has yet to confirm
    TopicArn:              topicARN,
  }

  result, err := subscribeTopic(context.TODO(), client, input)
  if err != nil {
    fmt.Println("Got an error subscribing to the topic:")
    fmt.Println(err)
    return
  }
  fmt.Println(*result.SubscriptionArn)
}

type SNSPublishAPI interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

func PublishMessage(c context.Context, api SNSPublishAPI, input *sns.PublishInput) (*sns.PublishOutput, error) {
	return api.Publish(c, input)
}

func PublishSMS(phoneNumber *string, firstName string){
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
      panic("configuration error, " + err.Error())
    }

    client := sns.NewFromConfig(cfg)

    msg := aws.String("Welcome to Picklejar, " + firstName + "!")
    //topicARN := aws.String("arn:aws:sns:us-east-1:375351057769:DeliveryNotification")

    input := &sns.PublishInput{
      PhoneNumber: phoneNumber,
      Message:  msg,

    }

    result, err := PublishMessage(context.TODO(), client, input)
    if err != nil {
      fmt.Println("Got an error publishing the message:")
      fmt.Println(err)
    return
    }

    fmt.Println("Message ID: " + *result.MessageId)
}
