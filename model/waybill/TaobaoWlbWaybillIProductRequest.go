package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家查询物流商产品类型接口 API请求
taobao.wlb.waybill.i.product

商家可以查询物流商的产品类型和服务能力。
*/
type TaobaoWlbWaybillIProductRequest struct {
    model.Params
    // 查询物流商电子面单产品类型入参
    waybillProductTypeRequest   *WaybillProductTypeRequest
}

// 初始化TaobaoWlbWaybillIProductRequest对象
func NewTaobaoWlbWaybillIProductRequest() *TaobaoWlbWaybillIProductRequest{
    return &TaobaoWlbWaybillIProductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIProductRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.product"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIProductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WaybillProductTypeRequest Setter
// 查询物流商电子面单产品类型入参
func (r *TaobaoWlbWaybillIProductRequest) SetWaybillProductTypeRequest(waybillProductTypeRequest *WaybillProductTypeRequest) error {
    r.waybillProductTypeRequest = waybillProductTypeRequest
    r.Set("waybill_product_type_request", waybillProductTypeRequest)
    return nil
}

// WaybillProductTypeRequest Getter
func (r TaobaoWlbWaybillIProductRequest) GetWaybillProductTypeRequest() *WaybillProductTypeRequest {
    return r.waybillProductTypeRequest
}
