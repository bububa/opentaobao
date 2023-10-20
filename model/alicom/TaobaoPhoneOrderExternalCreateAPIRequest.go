package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaophoneorderexternalcreateAPIRequest 数字虚拟话费外放下单接口 API请求
// taobao.phone.order.external.create
//
// 数字虚拟话费外放下单接口
type TaobaophoneorderexternalcreateAPIRequest struct {
	model.Params
	// 创建订单请求
	_createPhoneOrderReq *CreatePhoneOrderReq
}

// NewTaobaophoneorderexternalcreateRequest 初始化TaobaophoneorderexternalcreateAPIRequest对象
func NewTaobaophoneorderexternalcreateRequest() *TaobaophoneorderexternalcreateAPIRequest {
	return &TaobaophoneorderexternalcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaophoneorderexternalcreateAPIRequest) GetApiMethodName() string {
	return "taobao.phone.order.external.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaophoneorderexternalcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaophoneorderexternalcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreatePhoneOrderReq is CreatePhoneOrderReq Setter
// 创建订单请求
func (r *TaobaophoneorderexternalcreateAPIRequest) SetCreatePhoneOrderReq(_createPhoneOrderReq *CreatePhoneOrderReq) error {
	r._createPhoneOrderReq = _createPhoneOrderReq
	r.Set("create_phone_order_req", _createPhoneOrderReq)
	return nil
}

// GetCreatePhoneOrderReq CreatePhoneOrderReq Getter
func (r TaobaophoneorderexternalcreateAPIRequest) GetCreatePhoneOrderReq() *CreatePhoneOrderReq {
	return r._createPhoneOrderReq
}
