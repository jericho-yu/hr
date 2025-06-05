package config

// SystemConfig 系统配置
type SystemConfig struct {
	LogDir string `mapstructure:"log-dir" json:"log-dir" yaml:"log-dir"`
	Port   int    `mapstructure:"port" json:"port" yaml:"port"`
}
