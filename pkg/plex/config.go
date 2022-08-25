package plex

type Config struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}
