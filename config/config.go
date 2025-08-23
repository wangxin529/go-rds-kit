package config

type CacheType = string

var (
	CacheRedisType  CacheType = "redis"
	CacheMemoryType CacheType = "memory"
)

type Mysql struct {
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Logger   bool   `json:"logger"`
	DNS      string `json:"dns"`
}

type Memory struct {
	Type  CacheType `json:"type"`
	Redis *Redis    `json:"redis,optional"`
}
type Redis struct {
	Addr             []string `json:"addr"`
	Password         string   `json:"password"`
	MasterName       string   `json:"masterName"`
	SentinelPassword string   `json:"sentinelPassword"`
	DB               int      `json:"db"`
}
