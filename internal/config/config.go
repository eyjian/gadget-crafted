package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Cors int // 是否允许跨域（0 不允许，1 允许）
}
