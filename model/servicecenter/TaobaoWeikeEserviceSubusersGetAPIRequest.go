package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceSubusersGetAPIRequest 客服外包订单分配的商家子账号列表 API请求
// taobao.weike.eservice.subusers.get
//
// 获取客服外包订单分配的商家子账号列表，以及授权状态
type TaobaoWeikeEserviceSubusersGetAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// NewTaobaoWeikeEserviceSubusersGetRequest 初始化TaobaoWeikeEserviceSubusersGetAPIRequest对象
func NewTaobaoWeikeEserviceSubusersGetRequest() *TaobaoWeikeEserviceSubusersGetAPIRequest {
	return &TaobaoWeikeEserviceSubusersGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceSubusersGetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.subusers.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceSubusersGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *TaobaoWeikeEserviceSubusersGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoWeikeEserviceSubusersGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
