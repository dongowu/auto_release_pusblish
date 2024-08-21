/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2024/8/21
* Time: 11:45
* Software: GoLand
 */

package client

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	GeminiClient *genai.Client
)

const apiKey = "AIzaSyB4JZTqL0cIMqGAVduV9SkgNdLnePy0Xik"

func init() {
	client, err := genai.NewClient(context.TODO(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("failed to init  Gemini Client error: %s", err)
		return
	}

	GeminiClient = client
}
