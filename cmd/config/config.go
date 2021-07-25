package config

//config data

type Config struct {
    Protocol  string
    Host string
    Port string
    }

func New() *Config {
    return &Config{
        Protocol: "tcp",
        Host: "127.0.0.1",
        Port: ":8080",
    }
}
