package auth

import "github.com/kprince/weapp/v3/request"

// 用户信息
type Auth struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewAuth(request *request.Request, conbineURI func(url string, req interface{}, withToken bool) (string, error)) *Auth {
	sm := Auth{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
