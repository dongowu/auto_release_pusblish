/**
* Author: Jason
* Email: wudongdong@rongma.com
* Date: 2024/8/21
* Time: 11:36
* Software: GoLand
 */

package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/dongowu/auto_release_publish/geminipro_web/pkg/gemini"
)

func main() {

	h := server.Default(server.WithHostPorts(":8080"))

	fmt.Println(gemini.NewGemini(context.TODO()).GenerateContent(context.TODO(), "hello"))

	//h.Use()
	//
	//h.POST("/", func(c context.Context, ctx *app.RequestContext) {
	//
	//})
	h.Spin()

}
