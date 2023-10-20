package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbcrossborderwaybillgetAPIRequest 集货商家pdf和云打印面单获取，pdf需要配置白名单 API请求
// taobao.wlb.crossborder.waybill.get
//
// 【TOF】欧洲供应商PDF格式电子面单渲染下发
//
//	需求链接：https://aone.alibaba-inc.com/req/21210808
type TaobaowlbcrossborderwaybillgetAPIRequest struct {
	model.Params
	// 菜鸟物流单号
	_orderCode string
}

// NewTaobaowlbcrossborderwaybillgetRequest 初始化TaobaowlbcrossborderwaybillgetAPIRequest对象
func NewTaobaowlbcrossborderwaybillgetRequest() *TaobaowlbcrossborderwaybillgetAPIRequest {
	return &TaobaowlbcrossborderwaybillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbcrossborderwaybillgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.crossborder.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbcrossborderwaybillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbcrossborderwaybillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 菜鸟物流单号
func (r *TaobaowlbcrossborderwaybillgetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbcrossborderwaybillgetAPIRequest) GetOrderCode() string {
	return r._orderCode
}
