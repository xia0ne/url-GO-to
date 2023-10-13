package model

type Config struct {
	Config struct {
		BaseUrl string `yaml:"base_url"`
		Port    string `yaml:"port"`
	} `yaml:"config"`
	Redis struct {
		Addr          string `yaml:"Addr"`
		Passwd        string `yaml:"Passwd"`
		DB            string `yaml:"DB"`
		CacheDuration string `yaml:"CacheDuration"`
	} `yaml:"redis"`
}
