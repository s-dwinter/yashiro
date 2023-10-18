package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/s-dwinter/yashiro/pkg/config"
)

// Defines errors
var (
	ErrNotfoundValueConfig = errors.New("not found value config")
	ErrValueIsEmpty        = errors.New("value is empty")
)

// Client is the external stores client.
type Client interface {
	GetValues(ctx context.Context, ignoreEmpty bool) (Values, error)
}

// New returns a new Client.
func New(cfg *config.Config) (Client, error) {
	if cfg.Aws != nil {
		return newAwsClient(cfg.Aws)
	}

	return nil, ErrNotfoundValueConfig
}

// Values are stored values from external stores.
type Values map[string]any

// SetValue sets the getting value from external stores. If value is json string, is set
// as map[string]any.
func (v Values) SetValue(cfg config.Value, value *string) error {
	if value == nil || len(*value) == 0 {
		return ErrValueIsEmpty
	}

	if v == nil {
		v = make(Values)
	}

	var val any
	if cfg.GetIsJSON() {
		val = make(map[string]any)
		if err := json.Unmarshal([]byte(*value), &val); err != nil {
			return fmt.Errorf("%w: invalid json string", err)
		}
	} else {
		val = *value
	}

	v[cfg.GetReferenceName()] = val

	return nil
}
