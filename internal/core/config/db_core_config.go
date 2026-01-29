package config

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Name     string `env:"POSTGRES_DB"`
	SSLMode  string `env:"POSTGRES_SSLMODE"`
}


func NewPostgresConfig() *PostgresConfig {
	return &PostgresConfig{}
}

type RedisConfig struct {
	RedisHost     string `env:"REDIS_HOST"`
	RedisPort     string `env:"REDIS_PORT"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{}
}

type DBCoreConfig struct {
	DBPostgres PostgresConfig
	DBRedis    RedisConfig
}

func LoadDBCoreConfig() (*DBCoreConfig, error) {
	var cfg DBCoreConfig

	return &cfg, nil
}
