package config

// Config 总配置
type Config struct {
	System SystemConfig `mapstructure:"system" json:"system" yaml:"system"`
	Rezip  RezipConfig  `mapstructure:"rezip" json:"rezip" yaml:"rezip"`
}
