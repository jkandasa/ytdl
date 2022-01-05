package types

type Config struct {
	WebDirectory    string `yaml:"web_directory"`
	EnableProfiling bool   `yaml:"enable_profiling"`
	BindAddress     string `yaml:"bind_address"`
	Port            uint   `yaml:"port"`
}
