package idleisv

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleUserPermitQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleUserPermitQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
