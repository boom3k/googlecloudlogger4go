package googlecloudlogger4go

import (
	"cloud.google.com/go/logging"
	"context"
	"google.golang.org/api/option"
	"log"
)

func InitializeLogger(logName, projectID string, loggingKey []byte, serverity logging.Severity, ctx context.Context) *log.Logger {
	client, err := logging.NewClient(ctx, projectID, option.WithCredentialsJSON(loggingKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		panic(err)
	}
	log.Printf("Cloud Logger intialized...\nLogID: %s\nProjectID: %s\nSeverity: %s\n\n", logName, projectID, serverity)
	return client.Logger(logName).StandardLogger(serverity)
}
