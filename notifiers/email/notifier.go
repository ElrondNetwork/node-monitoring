package email

import (
	"fmt"
	"net/smtp"

	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/node-monitoring/config"
	"github.com/ElrondNetwork/node-monitoring/data"
)

var log = logger.GetOrCreate("eventNotifier")

// ArgsEmailNotifier defines the arguments needed to create a new email notifier
type ArgsEmailNotifier struct {
	Config *config.Email
}

type emailNotifier struct {
	config *config.Email
}

// NewEmailNotifier will create a new email notifier instance
func NewEmailNotifier(args ArgsEmailNotifier) (*emailNotifier, error) {
	en := &emailNotifier{
		config: args.Config,
	}

	return en, nil
}

// PushMessage will push the notification
func (en *emailNotifier) PushMessage(msg data.NotificationMessage) error {
	return en.push(msg)
}

func (en *emailNotifier) push(msg data.NotificationMessage) error {
	auth := smtp.PlainAuth("", en.config.EmailUsername, en.config.EmailPassword, en.config.EmailHost)

	smtpHost := fmt.Sprintf("%s:%d", en.config.EmailHost, en.config.EmailPort)

	err := smtp.SendMail(
		smtpHost,
		auth,
		en.config.From,
		en.config.To,
		en.msgToEmailMessageBytes(msg),
	)
	if err != nil {
		return err
	}

	log.Info("Email sent", "to", en.config.To)

	return nil
}

func (en *emailNotifier) msgToEmailMessageBytes(msg data.NotificationMessage) []byte {
	msgStr := fmt.Sprintf("To: %s\r\nSubject: %s!\r\n\r\n%s\r\n", "Email", "Nodes rating", msg.Message)
	return []byte(msgStr)
}

// GetID will return the identifier for email notifier
func (en *emailNotifier) GetID() string {
	return "SimpleEmail"
}

// IsInterfaceNil returns true if there is no value under the interface
func (en *emailNotifier) IsInterfaceNil() bool {
	return en == nil
}
