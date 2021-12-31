package googlecloudlogger4go

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"strings"
)

type GoogleCloudLogger struct {
	LogName           string
	ProjectID         string
	ServiceAccountKey []byte
	Context           context.Context
	Client            *logging.Client
	OutputStream      *logging.Logger
	Severity          logging.Severity
}

func Build(logName, projectID string, serviceAccountKey []byte, severity logging.Severity, ctx context.Context) *GoogleCloudLogger {
	client, err := logging.NewClient(ctx, projectID, option.WithCredentialsJSON(serviceAccountKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	googleCloudLogger := &GoogleCloudLogger{
		LogName:   logName,
		ProjectID: projectID,
		Client:    client,
		Context:   ctx,
		Severity:  severity,
	}
	googleCloudLogger.OutputStream = client.Logger(logName)
	return googleCloudLogger
}

func (receiver *GoogleCloudLogger) Printf(text string, v interface{}) {
	receiver.OutputStream.StandardLogger(receiver.Severity).Printf(text, v)
	receiver.OutputStream.Flush()
	log.Printf(strings.ToUpper(fmt.Sprint(receiver.Severity))+" - "+text, v)
}

func (receiver *GoogleCloudLogger) Println(text string) {
	receiver.OutputStream.StandardLogger(receiver.Severity).Println(text)
	receiver.OutputStream.Flush()
	log.Println(strings.ToUpper(fmt.Sprint(receiver.Severity)) + " - " + text)
}

//func main() {
//	gcl := Build("Default", "workspace-directory-tool", loggingKey, logging.Info, context.Background())
//	gcl.Println("INFO")
//	gcl.Severity = logging.Debug
//	gcl.Println("DEBUG")
//	gcl.Severity = logging.Default
//	gcl.Println("DEFAULT")
//	gcl.Severity = logging.Critical
//	gcl.Println("CRITICAL")
//	gcl.Severity = logging.Emergency
//	gcl.Println("EMERGENCY")
//	gcl.Severity = logging.Error
//	gcl.Println("ERROR")
//	gcl.Client.Ping(context.Background())
//}
