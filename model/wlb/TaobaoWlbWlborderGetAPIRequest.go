package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwlbordergetAPIRequest 根据物流宝订单编号查询物流宝订单概要信息 API请求
// taobao.wlb.wlborder.get
//
// 根据物流宝订单编号查询物流宝订单概要信息
type TaobaowlbwlbordergetAPIRequest struct {
	model.Params
	// 物流宝订单编码
	_wlbOrderCode string
}

// NewTaobaowlbwlbordergetRequest 初始化TaobaowlbwlbordergetAPIRequest对象
func NewTaobaowlbwlbordergetRequest() *TaobaowlbwlbordergetAPIRequest {
	return &TaobaowlbwlbordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwlbordergetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wlborder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwlbordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwlbordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWlbOrderCode is WlbOrderCode Setter
// 物流宝订单编码
func (r *TaobaowlbwlbordergetAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
	r._wlbOrderCode = _wlbOrderCode
	r.Set("wlb_order_code", _wlbOrderCode)
	return nil
}

// GetWlbOrderCode WlbOrderCode Getter
func (r TaobaowlbwlbordergetAPIRequest) GetWlbOrderCode() string {
	return r._wlbOrderCode
}
