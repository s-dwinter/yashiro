package yashiro_test

import (
	"context"
	"log"
	"os"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/s-dwinter/yashiro"
	"github.com/s-dwinter/yashiro/pkg/config"
)

func Example() {
	ctx := context.TODO()

	sdkConfig, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	refName := "example"
	cfg := &config.Config{
		Aws: &config.AwsConfig{
			ParameterStoreValues: []config.AwsParameterStoreValueConfig{
				{
					ValueConfig: config.ValueConfig{
						Name:   "/example",
						Ref:    &refName,
						IsJSON: true,
					},
				},
			},
			SdkConfig: &sdkConfig,
		},
	}

	eng, err := yashiro.NewEngine(cfg)
	if err != nil {
		log.Fatalf("failed to create engine: %s", err)
	}

	text := `This is example code. The message is {{ .example.message }}.`

	if err := eng.Render(ctx, text, os.Stdout); err != nil {
		log.Fatalf("failed to render text: %s", err)
	}
}
