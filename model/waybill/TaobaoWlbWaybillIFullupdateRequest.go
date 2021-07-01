package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
面单信息更新接口v1.0 API请求
taobao.wlb.waybill.i.fullupdate

商家更新电子面单号对应的订单信息。<br/><br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；<br/><br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
*/
type TaobaoWlbWaybillIFullupdateAPIRequest struct {
    model.Params
    // 更新面单信息请求
    _waybillApplyFullUpdateRequest   *WaybillApplyFullUpdateRequest
}

// 初始化TaobaoWlbWaybillIFullupdateAPIRequest对象
func NewTaobaoWlbWaybillIFullupdateRequest() *TaobaoWlbWaybillIFullupdateAPIRequest{
    return &TaobaoWlbWaybillIFullupdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillIFullupdateAPIRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.fullupdate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillIFullupdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WaybillApplyFullUpdateRequest Setter
// 更新面单信息请求
func (r *TaobaoWlbWaybillIFullupdateAPIRequest) SetWaybillApplyFullUpdateRequest(_waybillApplyFullUpdateRequest *WaybillApplyFullUpdateRequest) error {
    r._waybillApplyFullUpdateRequest = _waybillApplyFullUpdateRequest
    r.Set("waybill_apply_full_update_request", _waybillApplyFullUpdateRequest)
    return nil
}

// WaybillApplyFullUpdateRequest Getter
func (r TaobaoWlbWaybillIFullupdateAPIRequest) GetWaybillApplyFullUpdateRequest() *WaybillApplyFullUpdateRequest {
    return r._waybillApplyFullUpdateRequest
}
