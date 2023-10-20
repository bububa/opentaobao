package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaophoneorderexternalstatusAPIRequest 话费外放订单状态接口 API请求
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
type TaobaophoneorderexternalstatusAPIRequest struct {
	model.Params
	// 状态查询请求
	_queryOrderReq *QueryOrderReq
}

// NewTaobaophoneorderexternalstatusRequest 初始化TaobaophoneorderexternalstatusAPIRequest对象
func NewTaobaophoneorderexternalstatusRequest() *TaobaophoneorderexternalstatusAPIRequest {
	return &TaobaophoneorderexternalstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaophoneorderexternalstatusAPIRequest) GetApiMethodName() string {
	return "taobao.phone.order.external.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaophoneorderexternalstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaophoneorderexternalstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryOrderReq is QueryOrderReq Setter
// 状态查询请求
func (r *TaobaophoneorderexternalstatusAPIRequest) SetQueryOrderReq(_queryOrderReq *QueryOrderReq) error {
	r._queryOrderReq = _queryOrderReq
	r.Set("query_order_req", _queryOrderReq)
	return nil
}

// GetQueryOrderReq QueryOrderReq Getter
func (r TaobaophoneorderexternalstatusAPIRequest) GetQueryOrderReq() *QueryOrderReq {
	return r._queryOrderReq
}
