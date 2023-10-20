package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvUserQueryAPIRequest 服务商ISV闲鱼用户信息查询 API请求
// alibaba.idle.isv.user.query
//
// 服务商ISV闲鱼用户信息查询
type AlibabaIdleIsvUserQueryAPIRequest struct {
	model.Params
}

// NewAlibabaIdleIsvUserQueryRequest 初始化AlibabaIdleIsvUserQueryAPIRequest对象
func NewAlibabaIdleIsvUserQueryRequest() *AlibabaIdleIsvUserQueryAPIRequest {
	return &AlibabaIdleIsvUserQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvUserQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvUserQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.user.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvUserQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvUserQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaIdleIsvUserQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvUserQueryRequest()
	},
}

// GetAlibabaIdleIsvUserQueryRequest 从 sync.Pool 获取 AlibabaIdleIsvUserQueryAPIRequest
func GetAlibabaIdleIsvUserQueryAPIRequest() *AlibabaIdleIsvUserQueryAPIRequest {
	return poolAlibabaIdleIsvUserQueryAPIRequest.Get().(*AlibabaIdleIsvUserQueryAPIRequest)
}

// ReleaseAlibabaIdleIsvUserQueryAPIRequest 将 AlibabaIdleIsvUserQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvUserQueryAPIRequest(v *AlibabaIdleIsvUserQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvUserQueryAPIRequest.Put(v)
}
