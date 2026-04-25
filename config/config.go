package config

type Config struct {
	Key string
}

/*func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр key в переменные окружения")
	}
	return &Config{
		Key: key,
	}
}
*/

func (c *Config) GetAPIKey() string {
	return c.Key
}
