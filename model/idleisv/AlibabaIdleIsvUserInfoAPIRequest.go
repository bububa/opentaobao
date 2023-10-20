package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvUserInfoAPIRequest 闲鱼用户信息查询接口 API请求
// alibaba.idle.isv.user.info
//
// 闲鱼用户信息查询接口
type AlibabaIdleIsvUserInfoAPIRequest struct {
	model.Params
}

// NewAlibabaIdleIsvUserInfoRequest 初始化AlibabaIdleIsvUserInfoAPIRequest对象
func NewAlibabaIdleIsvUserInfoRequest() *AlibabaIdleIsvUserInfoAPIRequest {
	return &AlibabaIdleIsvUserInfoAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvUserInfoAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvUserInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.user.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvUserInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvUserInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaIdleIsvUserInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvUserInfoRequest()
	},
}

// GetAlibabaIdleIsvUserInfoRequest 从 sync.Pool 获取 AlibabaIdleIsvUserInfoAPIRequest
func GetAlibabaIdleIsvUserInfoAPIRequest() *AlibabaIdleIsvUserInfoAPIRequest {
	return poolAlibabaIdleIsvUserInfoAPIRequest.Get().(*AlibabaIdleIsvUserInfoAPIRequest)
}

// ReleaseAlibabaIdleIsvUserInfoAPIRequest 将 AlibabaIdleIsvUserInfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvUserInfoAPIRequest(v *AlibabaIdleIsvUserInfoAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvUserInfoAPIRequest.Put(v)
}
