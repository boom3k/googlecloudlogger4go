package googlecloudlogger4go

import (
	"cloud.google.com/go/logging"
	"context"
	"google.golang.org/api/option"
	"log"
)

type CloudLogger struct {
	Logger *log.Logger
}

func InitializeLogger(projectID, logID string, loggingKey []byte, logType logging.Severity, ctx context.Context) *CloudLogger {
	client, err := logging.NewClient(ctx, projectID, option.WithCredentialsJSON(loggingKey))
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	return &CloudLogger{Logger: client.Logger(logID).StandardLogger(logType)}
}

func (receiver CloudLogger) Log(text string, v ...interface{}) {
	log.Printf(text, v)
}
