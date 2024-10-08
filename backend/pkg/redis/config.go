package redis

type (
	Config struct {
		Host string
		Port string
	}
	PubSubConfig struct {
		Host string
		Port string
	}
	PoolConfig struct {
		Host    string
		Port    string
		MaxIdle int
	}
)
