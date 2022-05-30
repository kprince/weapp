package security

import (
	"github.com/kprince/weapp/v3/request"
)

type Security struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewSecurity(request *request.Request, conbineURI func(url string, req interface{}, withToken bool) (string, error)) *Security {
	sm := Security{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
