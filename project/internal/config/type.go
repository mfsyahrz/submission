package config

type Config struct {
	Service  Service
	Postgres Postgres
}

type Service struct {
	Name string `env:"SERVICE_NAME,default=sb_project"`
	Port Port
}

type Port struct {
	REST string `env:"SERVICE_PORT_REST,default=8080"`
	GRPC string `env:"SERVICE_PORT_GRPC,default=8081"`
}

type Postgres struct {
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	Name            string `env:"POSTGRES_NAME,required"`
	Port            string `env:"POSTGRES_PORT,default=5432"`
	Host            string `env:"POSTGRES_HOST,default=localhost"`
	MaxOpenConns    string `env:"POSTGRES_MAX_OPEN_CONNS,default=5"`
	MaxConnLifetime string `env:"POSTGRES_MAX_CONN,default=10m"`
	MaxIdleLifetime string `env:"POSTGRES_MAX_IDLE,default=5m"`
}
