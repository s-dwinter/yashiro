package yashiro

import (
	"github.com/s-dwinter/yashiro/pkg/config"
	"github.com/s-dwinter/yashiro/pkg/engine"
)

// Engine initializes external store client and template.
type Engine = engine.Engine

var (
	// NewConfigFromFile returns a new Config from file.
	NewConfigFromFile = config.NewFromFile

	// NewEngine returns a new Engine.
	NewEngine = engine.New
)
