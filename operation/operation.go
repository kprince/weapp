package operation

import (
	"github.com/kprince/weapp/v3/request"
)

type Operation struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewOperation(request *request.Request, conbineURI func(url string, req interface{}, withToken bool) (string, error)) *Operation {
	sm := Operation{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
