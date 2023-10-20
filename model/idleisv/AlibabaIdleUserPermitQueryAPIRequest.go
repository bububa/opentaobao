package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleuserpermitqueryAPIRequest 查询服务商与卖家之间的订单消息绑定关系 API请求
// alibaba.idle.user.permit.query
//
// 查询服务商与卖家之间的订单消息绑定关系
type AlibabaidleuserpermitqueryAPIRequest struct {
	model.Params
}

// NewAlibabaidleuserpermitqueryRequest 初始化AlibabaidleuserpermitqueryAPIRequest对象
func NewAlibabaidleuserpermitqueryRequest() *AlibabaidleuserpermitqueryAPIRequest {
	return &AlibabaidleuserpermitqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleuserpermitqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.user.permit.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleuserpermitqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleuserpermitqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
