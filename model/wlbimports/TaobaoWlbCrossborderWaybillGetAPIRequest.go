package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbCrossborderWaybillGetAPIRequest 集货商家pdf和云打印面单获取，pdf需要配置白名单 API请求
// taobao.wlb.crossborder.waybill.get
//
// 【TOF】欧洲供应商PDF格式电子面单渲染下发
//  需求链接：https://aone.alibaba-inc.com/req/21210808
type TaobaoWlbCrossborderWaybillGetAPIRequest struct {
	model.Params
	// 菜鸟物流单号
	_orderCode string
}

// NewTaobaoWlbCrossborderWaybillGetRequest 初始化TaobaoWlbCrossborderWaybillGetAPIRequest对象
func NewTaobaoWlbCrossborderWaybillGetRequest() *TaobaoWlbCrossborderWaybillGetAPIRequest {
	return &TaobaoWlbCrossborderWaybillGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbCrossborderWaybillGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.crossborder.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbCrossborderWaybillGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderCode Setter
// 菜鸟物流单号
func (r *TaobaoWlbCrossborderWaybillGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// Get OrderCode Getter
func (r TaobaoWlbCrossborderWaybillGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}
