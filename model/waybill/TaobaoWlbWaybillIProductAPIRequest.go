package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwaybilliproductAPIRequest 商家查询物流商产品类型接口 API请求
// taobao.wlb.waybill.i.product
//
// 商家可以查询物流商的产品类型和服务能力。
type TaobaowlbwaybilliproductAPIRequest struct {
	model.Params
	// 查询物流商电子面单产品类型入参
	_waybillProductTypeRequest *WaybillProductTypeRequest
}

// NewTaobaowlbwaybilliproductRequest 初始化TaobaowlbwaybilliproductAPIRequest对象
func NewTaobaowlbwaybilliproductRequest() *TaobaowlbwaybilliproductAPIRequest {
	return &TaobaowlbwaybilliproductAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwaybilliproductAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.product"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwaybilliproductAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwaybilliproductAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWaybillProductTypeRequest is WaybillProductTypeRequest Setter
// 查询物流商电子面单产品类型入参
func (r *TaobaowlbwaybilliproductAPIRequest) SetWaybillProductTypeRequest(_waybillProductTypeRequest *WaybillProductTypeRequest) error {
	r._waybillProductTypeRequest = _waybillProductTypeRequest
	r.Set("waybill_product_type_request", _waybillProductTypeRequest)
	return nil
}

// GetWaybillProductTypeRequest WaybillProductTypeRequest Getter
func (r TaobaowlbwaybilliproductAPIRequest) GetWaybillProductTypeRequest() *WaybillProductTypeRequest {
	return r._waybillProductTypeRequest
}
