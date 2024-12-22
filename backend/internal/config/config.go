package config

type Config struct {
    DatabaseURL string
    Port        string
    JWTSecret   string
}

func LoadConfig() (*Config, error) {
    // TODO: Implementar carregamento de configurações do .env
    return &Config{
        Port: "8080",
    }, nil
} 