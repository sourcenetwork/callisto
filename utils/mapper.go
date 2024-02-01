package utils

import (
	"cosmossdk.io/simapp/params"
	junoparams "github.com/forbole/juno/v5/types/params"
)

// ToJunoEncodingConfig converts an SDK EncodingConfig struct into a juno EncodingConfig struct
func ToJunoEncodingConfig(cfg params.EncodingConfig) junoparams.EncodingConfig {
	return junoparams.EncodingConfig{
		InterfaceRegistry: cfg.InterfaceRegistry,
		Codec:             cfg.Codec,
		TxConfig:          cfg.TxConfig,
		Amino:             cfg.Amino,
	}
}

// FromJunoEncodingConfig converts a juno EncodingConfig struct into a SDK EncodingConfig struct
func FromJunoEncodingConfig(cfg junoparams.EncodingConfig) params.EncodingConfig {
	return params.EncodingConfig{
		InterfaceRegistry: cfg.InterfaceRegistry,
		Codec:             cfg.Codec,
		TxConfig:          cfg.TxConfig,
		Amino:             cfg.Amino,
	}
}
