package conf

import "user_server/pkg/utils"

type ServerConfig struct {
	Addr string
}

func (r *ServerConfig) Read(name string) {
	data := utils.Getenv(name, `1`)
	r.Addr = data
}
