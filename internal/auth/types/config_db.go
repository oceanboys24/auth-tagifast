package types

type ConfigDB struct {
	DBHost     string `env:"SUPABASE_URL"`
	DBPort     string `env:"SUPABASE_PORT"`
	DBUser     string `env:"SUPABASE_USER"`
	DBName     string `env:"SUPABASE_NAME"`
	DBPassword string `env:"SUPABASE_PASSWORD"`
}
