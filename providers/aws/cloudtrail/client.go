package cloudtrail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/cloudquery/cloudquery/providers/aws/resource"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Client struct {
	session   *session.Session
	db        *gorm.DB
	log       *zap.Logger
	accountID string
	region    string
	svc       *cloudtrail.CloudTrail
}

func NewClient(session *session.Session, awsConfig *aws.Config, db *gorm.DB, log *zap.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		session:   session,
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       cloudtrail.New(session, awsConfig),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "trails":
		return c.trails(config)
	default:
		return fmt.Errorf("unsupported resource autoscaling.%s", resource)
	}
}
