package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWaybillIProductAPIRequest 商家查询物流商产品类型接口 API请求
// taobao.wlb.waybill.i.product
//
// 商家可以查询物流商的产品类型和服务能力。
type TaobaoWlbWaybillIProductAPIRequest struct {
	model.Params
	// 查询物流商电子面单产品类型入参
	_waybillProductTypeRequest *WaybillProductTypeRequest
}

// NewTaobaoWlbWaybillIProductRequest 初始化TaobaoWlbWaybillIProductAPIRequest对象
func NewTaobaoWlbWaybillIProductRequest() *TaobaoWlbWaybillIProductAPIRequest {
	return &TaobaoWlbWaybillIProductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIProductAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIProductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWaybillIProductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillProductTypeRequest is WaybillProductTypeRequest Setter
// 查询物流商电子面单产品类型入参
func (r *TaobaoWlbWaybillIProductAPIRequest) SetWaybillProductTypeRequest(_waybillProductTypeRequest *WaybillProductTypeRequest) error {
	r._waybillProductTypeRequest = _waybillProductTypeRequest
	r.Set("waybill_product_type_request", _waybillProductTypeRequest)
	return nil
}

// GetWaybillProductTypeRequest WaybillProductTypeRequest Getter
func (r TaobaoWlbWaybillIProductAPIRequest) GetWaybillProductTypeRequest() *WaybillProductTypeRequest {
	return r._waybillProductTypeRequest
}
