package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
面单信息更新接口v1.0 APIRequest
taobao.wlb.waybill.i.fullupdate

商家更新电子面单号对应的订单信息。<br/><br/>a.涉及到订单信息（如拆/合单、取消订单等）、修改发货地的逻辑时候，需要使用CANCEL+GET组合操作。一般的面单信息更新使用该接口；<br/><br/>b.该接口是全量修改接口，对于不做更新的字段也要把原有的字段值传进去，否则做为修改为空处理。
*/
type TaobaoWlbWaybillIFullupdateRequest struct {
    model.Params

    // 更新面单信息请求
    waybillApplyFullUpdateRequest   *WaybillApplyFullUpdateRequest 

}

func NewTaobaoWlbWaybillIFullupdateRequest() *TaobaoWlbWaybillIFullupdateRequest{
    return &TaobaoWlbWaybillIFullupdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWaybillIFullupdateRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.fullupdate"
}

func (r TaobaoWlbWaybillIFullupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWaybillIFullupdateRequest) SetWaybillApplyFullUpdateRequest(waybillApplyFullUpdateRequest *WaybillApplyFullUpdateRequest) error {
    r.waybillApplyFullUpdateRequest = waybillApplyFullUpdateRequest
    r.Set("waybill_apply_full_update_request", waybillApplyFullUpdateRequest)
    return nil
}

func (r TaobaoWlbWaybillIFullupdateRequest) GetWaybillApplyFullUpdateRequest() *WaybillApplyFullUpdateRequest {
    return r.waybillApplyFullUpdateRequest
}

