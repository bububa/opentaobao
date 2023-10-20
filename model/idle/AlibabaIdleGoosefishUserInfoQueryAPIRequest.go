package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleGoosefishUserInfoQueryAPIRequest 闲鱼三方容器用户信息获取 API请求
// alibaba.idle.goosefish.user.info.query
//
// 闲鱼三方容器用户信息获取
type AlibabaIdleGoosefishUserInfoQueryAPIRequest struct {
	model.Params
}

// NewAlibabaIdleGoosefishUserInfoQueryRequest 初始化AlibabaIdleGoosefishUserInfoQueryAPIRequest对象
func NewAlibabaIdleGoosefishUserInfoQueryRequest() *AlibabaIdleGoosefishUserInfoQueryAPIRequest {
	return &AlibabaIdleGoosefishUserInfoQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleGoosefishUserInfoQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleGoosefishUserInfoQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.goosefish.user.info.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleGoosefishUserInfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleGoosefishUserInfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaIdleGoosefishUserInfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleGoosefishUserInfoQueryRequest()
	},
}

// GetAlibabaIdleGoosefishUserInfoQueryRequest 从 sync.Pool 获取 AlibabaIdleGoosefishUserInfoQueryAPIRequest
func GetAlibabaIdleGoosefishUserInfoQueryAPIRequest() *AlibabaIdleGoosefishUserInfoQueryAPIRequest {
	return poolAlibabaIdleGoosefishUserInfoQueryAPIRequest.Get().(*AlibabaIdleGoosefishUserInfoQueryAPIRequest)
}

// ReleaseAlibabaIdleGoosefishUserInfoQueryAPIRequest 将 AlibabaIdleGoosefishUserInfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleGoosefishUserInfoQueryAPIRequest(v *AlibabaIdleGoosefishUserInfoQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleGoosefishUserInfoQueryAPIRequest.Put(v)
}
