package googlepicker

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// Config is the configuration for the Google Picker plugin.
type Config struct {
}

// plugin is the implementation of the Google Picker plugin.
type plugin struct {
	config Config
}

var _ blazarplugin.Plugin = (*plugin)(nil)

// NewPlugin creates a new Google Picker plugin.
func NewPlugin(config Config) blazarplugin.Plugin {
	return &plugin{
		config: config,
	}
}

// Register registers the plugin against the go-app handler and the HTTP mux.
func (p *plugin) Register(handler *app.Handler, mux *http.ServeMux) {
	slog.InfoContext(context.TODO(), "Registering Google Picker plugin", "config", p.config)

	handler.Scripts = append(handler.Scripts,
		"https://apis.google.com/js/api.js",
		"https://accounts.google.com/gsi/client",
	)

	handler.CacheableResources = append(handler.CacheableResources,
		"https://apis.google.com/js/api.js",
		"https://accounts.google.com/gsi/client",
	)
}
