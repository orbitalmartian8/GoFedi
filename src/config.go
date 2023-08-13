package main

type Config struct {
    InstanceURL   string
    ClientID      string
    ClientSecret  string
    AccessToken  string
}

func LoadConfig() *Config {
    return &Config{
        InstanceURL:  "https://linuxrocks.online",
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        AccessToken:  "your-access-token",
    }
}
