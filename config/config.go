package config

type Config struct {
	Redis struct {
		RedisAddr      string `yaml:"redis_addr"`
		RedisDB        string `yaml:"redis_db"`
		RedisDefaultDB string `yaml:"redis_default_db"`
		MinIdleConn    int    `yaml:"min_idle_conn"`
		PoolSize       int    `yaml:"pool_size"`
		PoolTimeout    int    `yaml:"pool_timeout"`
		Password       string `yaml:"password"`
		DB             int    `yaml:"db"`
	} `yaml:"redis_db"`
	Server struct {
		Host string `yaml:"host"`
	} `yaml:"server"`
}
