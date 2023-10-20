package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbordercancelAPIRequest 取消物流宝订单 API请求
// taobao.wlb.order.cancel
//
// 取消物流宝订单
type TaobaowlbordercancelAPIRequest struct {
	model.Params
	// 物流宝订单编号
	_wlbOrderCode string
}

// NewTaobaowlbordercancelRequest 初始化TaobaowlbordercancelAPIRequest对象
func NewTaobaowlbordercancelRequest() *TaobaowlbordercancelAPIRequest {
	return &TaobaowlbordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbordercancelAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWlbOrderCode is WlbOrderCode Setter
// 物流宝订单编号
func (r *TaobaowlbordercancelAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
	r._wlbOrderCode = _wlbOrderCode
	r.Set("wlb_order_code", _wlbOrderCode)
	return nil
}

// GetWlbOrderCode WlbOrderCode Getter
func (r TaobaowlbordercancelAPIRequest) GetWlbOrderCode() string {
	return r._wlbOrderCode
}
