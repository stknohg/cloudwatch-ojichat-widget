package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/greymd/ojichat/generator"
)

type MyEvent struct {
	Name     string `json:"name"`
	Describe bool   `json:"describe"`
	// widgetContext の定義は省略
}

func renderUsage() (string, error) {
	return strings.ReplaceAll(`## ojichat custom widget
おじさんがあいさつをしてくれます。

### パラメーター
以下のJSONを渡してください。

{tick}{tick}{tick}json
{
	"name": "<好きな名前>"
}
{tick}{tick}{tick}
`, "{tick}", "`"), nil
}

func renderOutput(message string) (string, error) {
	return strings.ReplaceAll(`<style>
.ojichat-license {
	float: right;
}
</style>
<pre>{message}</pre>
<div class="ojichat-license">
This widget uses 😘 <a href="https://github.com/greymd/ojichat" target="_blank">ojichat</a> 💦.
</div>`, "{message}", message), nil
}

func handler(ctx context.Context, event MyEvent) (string, error) {
	// ヘルプドキュメントを表示する場合は Describe がセットされてLambdaが呼び出される
	if event.Describe {
		return renderUsage()
	}

	// あいさつを返す
	config := generator.Config{
		TargetName:       event.Name,
		EmojiNum:         4,
		PunctuationLevel: 0,
	}
	message, err := generator.Start(config)
	if err != nil {
		return err.Error(), err
	}
	return renderOutput(message)
}

func main() {
	lambda.Start(handler)
}
