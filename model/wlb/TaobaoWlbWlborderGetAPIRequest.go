package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWlborderGetAPIRequest 根据物流宝订单编号查询物流宝订单概要信息 API请求
// taobao.wlb.wlborder.get
//
// 根据物流宝订单编号查询物流宝订单概要信息
type TaobaoWlbWlborderGetAPIRequest struct {
	model.Params
	// 物流宝订单编码
	_wlbOrderCode string
}

// NewTaobaoWlbWlborderGetRequest 初始化TaobaoWlbWlborderGetAPIRequest对象
func NewTaobaoWlbWlborderGetRequest() *TaobaoWlbWlborderGetAPIRequest {
	return &TaobaoWlbWlborderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWlborderGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wlborder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWlborderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWlbOrderCode is WlbOrderCode Setter
// 物流宝订单编码
func (r *TaobaoWlbWlborderGetAPIRequest) SetWlbOrderCode(_wlbOrderCode string) error {
	r._wlbOrderCode = _wlbOrderCode
	r.Set("wlb_order_code", _wlbOrderCode)
	return nil
}

// GetWlbOrderCode WlbOrderCode Getter
func (r TaobaoWlbWlborderGetAPIRequest) GetWlbOrderCode() string {
	return r._wlbOrderCode
}
