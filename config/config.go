package config

type Config struct {
	Port        int         `json:"port"`
	Redis       *Redis      `json:"redis"`
	Application Application `json:"application"`
}
type Application struct {
	StoragePath string `json:"storage_path"`
}
type Redis struct {
	Addr             []string `json:"addr"`
	Password         string   `json:"password"`
	Port             string   `json:"port"`
	MasterName       string   `json:"masterName"`
	SentinelPassword string   `json:"sentinelPassword"`
	DB               int      `json:"db"`
}
