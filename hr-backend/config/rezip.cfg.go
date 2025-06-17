package config

// RezipConfig 重复压缩配置
type RezipConfig struct {
	OutDir string `mapstructure:"out-dir" json:"out-dir" yaml:"out-dir"`
}
