package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPhoneOrderExternalCreateAPIRequest 数字虚拟话费外放下单接口 API请求
// taobao.phone.order.external.create
//
// 数字虚拟话费外放下单接口
type TaobaoPhoneOrderExternalCreateAPIRequest struct {
	model.Params
	// 创建订单请求
	_createPhoneOrderReq *CreatePhoneOrderReq
}

// NewTaobaoPhoneOrderExternalCreateRequest 初始化TaobaoPhoneOrderExternalCreateAPIRequest对象
func NewTaobaoPhoneOrderExternalCreateRequest() *TaobaoPhoneOrderExternalCreateAPIRequest {
	return &TaobaoPhoneOrderExternalCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPhoneOrderExternalCreateAPIRequest) GetApiMethodName() string {
	return "taobao.phone.order.external.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPhoneOrderExternalCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPhoneOrderExternalCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreatePhoneOrderReq is CreatePhoneOrderReq Setter
// 创建订单请求
func (r *TaobaoPhoneOrderExternalCreateAPIRequest) SetCreatePhoneOrderReq(_createPhoneOrderReq *CreatePhoneOrderReq) error {
	r._createPhoneOrderReq = _createPhoneOrderReq
	r.Set("create_phone_order_req", _createPhoneOrderReq)
	return nil
}

// GetCreatePhoneOrderReq CreatePhoneOrderReq Getter
func (r TaobaoPhoneOrderExternalCreateAPIRequest) GetCreatePhoneOrderReq() *CreatePhoneOrderReq {
	return r._createPhoneOrderReq
}
