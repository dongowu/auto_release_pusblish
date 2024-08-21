/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2024/8/21
* Time: 11:40
* Software: GoLand
 */

package gemini

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/appengine/log"

	"github.com/dongowu/auto_release_publish/geminipro_web/cmd/config"
	"github.com/dongowu/auto_release_publish/geminipro_web/internal/client"
)

type geminiPro struct {
	client *genai.Client
}

func NewGemini(ctx context.Context) *geminiPro {

	return &geminiPro{client: client.GeminiClient}
}

func (g *geminiPro) GenerateContent(ctx context.Context, input interface{}) (string, error) {

	model := g.client.GenerativeModel(config.Gemini15Flash)

	resp, err := model.GenerateContent(ctx, genai.Text("tell me who are you "))
	if err != nil {
		return "", err
	}

	log.Infof(ctx, "input :%s ,resp :%+v", input, resp)

	return string(resp.Candidates[0].Content.Parts[0].(genai.Text)), nil
}

func (g *geminiPro) GenerateStreamContext(ctx context.Context, input interface{}) (string, error) {

	model := g.client.GenerativeModel(config.Gemini15Flash)

	streamResp := model.GenerateContentStream(ctx, genai.Text("tell me who are you "))

	return string(streamResp.MergedResponse().Candidates[0].Content.Parts[0].(genai.Text)), nil

}

func (g *geminiPro) name() {

}
