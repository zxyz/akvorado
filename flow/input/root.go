package input

import (
	"akvorado/daemon"
	"akvorado/flow/decoder"
	"akvorado/reporter"
)

// Input is the interface any input should meet
type Input interface {
	// Start instructs an input to start producing flows on the returned channel.
	Start() (<-chan []*decoder.FlowMessage, error)
	// Stop instructs the input to stop producing flows.
	Stop() error
}

// Configuration the interface for the configuration for an input module.
type Configuration interface {
	// New instantiantes a new input from its configuration.
	New(r *reporter.Reporter, daemon daemon.Component, dec decoder.Decoder) (Input, error)
}
