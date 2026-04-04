package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/csmith/envflag/v2"
	"github.com/greboid/irc-bot/v5/plugins"
	"github.com/greboid/irc-bot/v5/rpc"
)

var (
	rpcHost        = flag.String("rpc-host", "localhost", "gRPC server to connect to")
	rpcPort        = flag.Int("rpc-port", 8001, "gRPC server port")
	rpcToken       = flag.String("rpc-token", "", "gRPC authentication token")
	channel        = flag.String("channel", "", "Channel to send messages to")
	webhookSecret  = flag.String("webhook-secret", "", "Secret for verifying webhook signatures")
	messagePrefix  = flag.String("message-prefix", "", "Prefix to add to the start of each message")
)

type notification struct {
	Image  string `json:"image"`
	Target string `json:"target"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func main() {
	envflag.Parse()

	helper, err := plugins.NewHelper(fmt.Sprintf("%s:%d", *rpcHost, uint16(*rpcPort)), *rpcToken)
	if err != nil {
		log.Fatalf("Unable to create plugin helper: %v", err)
	}

	if err := helper.RegisterWebhook("adze", webhookHandler(helper)); err != nil {
		log.Fatalf("Error registering webhook: %v", err)
	}
}

func webhookHandler(helper *plugins.PluginHelper) func(request *rpc.HttpRequest) *rpc.HttpResponse {
	return func(request *rpc.HttpRequest) *rpc.HttpResponse {
		if !verifySignature(request) {
			return &rpc.HttpResponse{Status: http.StatusForbidden}
		}

		n := notification{}
		if err := json.Unmarshal(request.Body, &n); err != nil {
			log.Printf("Failed to unmarshal webhook body: %v", err)
			return &rpc.HttpResponse{Status: http.StatusBadRequest}
		}

		msg := formatMessage(n)
		if err := helper.SendChannelMessage(*channel, msg); err != nil {
			log.Printf("Failed to send channel message: %v", err)
		}

		return &rpc.HttpResponse{Status: http.StatusNoContent}
	}
}

func verifySignature(request *rpc.HttpRequest) bool {
	var sigHeader string
	for i := range request.Header {
		if request.Header[i].Key == "X-Adze-Signature" {
			sigHeader = request.Header[i].Value
			break
		}
	}

	if sigHeader == "" {
		return false
	}

	if !strings.HasPrefix(sigHeader, "sha256=") {
		return false
	}

	mac := hmac.New(sha256.New, []byte(*webhookSecret))
	mac.Write(request.Body)
	expected := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(sigHeader), []byte(expected))
}

func formatMessage(n notification) string {
	var msg strings.Builder

	if *messagePrefix != "" {
		msg.WriteString(*messagePrefix)
		msg.WriteString(" ")
	}

	msg.WriteString(n.Target)
	msg.WriteString(": ")
	msg.WriteString(n.Status)
	msg.WriteString(" updating ")
	msg.WriteString(n.Image)

	if n.Status == "failure" && n.Error != "" {
		msg.WriteString(" — ")
		msg.WriteString(n.Error)
	}

	return msg.String()
}
