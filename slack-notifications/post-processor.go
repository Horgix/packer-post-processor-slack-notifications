package slacknotifications

import (
	"log"

	//awscommon "github.com/hashicorp/packer/builder/amazon/common"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	Channel string `mapstructure:"channel"`
	Webhook string `mapstructure:"webhook"`

	ctx interpolate.Context
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	log.Println("Decoding configuration...")
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {
	ui.Say("Processing...")
	log.Println(artifact.String())
	ui.Say("Sending Slack notification...")
	payload := slack.Payload{
		Text:    artifact.String(),
		Channel: p.config.Channel,
	}
	ui.Say("test")
	err := slack.Send(p.config.Webhook, "", payload)
	if len(err) > 0 {
		log.Println("error: %s\n", err)
	}
	return artifact, true, nil
}
