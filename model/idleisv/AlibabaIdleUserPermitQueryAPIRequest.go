package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleUserPermitQueryAPIRequest 查询服务商与卖家之间的订单消息绑定关系 API请求
// alibaba.idle.user.permit.query
//
// 查询服务商与卖家之间的订单消息绑定关系
type AlibabaIdleUserPermitQueryAPIRequest struct {
	model.Params
}

// NewAlibabaIdleUserPermitQueryRequest 初始化AlibabaIdleUserPermitQueryAPIRequest对象
func NewAlibabaIdleUserPermitQueryRequest() *AlibabaIdleUserPermitQueryAPIRequest {
	return &AlibabaIdleUserPermitQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleUserPermitQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleUserPermitQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleUserPermitQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleUserPermitQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaIdleUserPermitQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleUserPermitQueryRequest()
	},
}

// GetAlibabaIdleUserPermitQueryRequest 从 sync.Pool 获取 AlibabaIdleUserPermitQueryAPIRequest
func GetAlibabaIdleUserPermitQueryAPIRequest() *AlibabaIdleUserPermitQueryAPIRequest {
	return poolAlibabaIdleUserPermitQueryAPIRequest.Get().(*AlibabaIdleUserPermitQueryAPIRequest)
}

// ReleaseAlibabaIdleUserPermitQueryAPIRequest 将 AlibabaIdleUserPermitQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleUserPermitQueryAPIRequest(v *AlibabaIdleUserPermitQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleUserPermitQueryAPIRequest.Put(v)
}
