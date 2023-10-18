package engine

import (
	"context"
	"io"
	"text/template"

	"github.com/s-dwinter/yashiro/internal/client"
	"github.com/s-dwinter/yashiro/pkg/config"
)

type Engine interface {
	Render(ctx context.Context, text string, dest io.Writer) error
}

type engine struct {
	client   client.Client
	template *template.Template
	option   *engineOpts
}

func New(cfg *config.Config, option ...EngineOption) (Engine, error) {
	opts := defaultEngineOpts
	for _, o := range option {
		o(opts)
	}

	cli, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	tmpl := template.New("yashiro").Option("missingkey=error").Funcs(funcMap())

	return &engine{
		client:   cli,
		template: tmpl,
		option:   opts,
	}, nil
}

func (e engine) Render(ctx context.Context, text string, dest io.Writer) error {
	values, err := e.client.GetValues(ctx, e.option.IgnoreEmpty)
	if err != nil {
		return err
	}

	return e.render(text, dest, values)
}

func (e engine) render(text string, dest io.Writer, data any) error {
	if _, err := e.template.Parse(text); err != nil {
		return err
	}

	return e.template.Execute(dest, data)
}
