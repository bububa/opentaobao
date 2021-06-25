package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取物流服务商电子面单号v1.0 APIRequest
taobao.wlb.waybill.i.get

商家根据订单信息，实时、批量获取指定物流服务商的电子面单号。
*/
type TaobaoWlbWaybillIGetRequest struct {
    model.Params

    // 面单申请
    waybillApplyNewRequest   *WaybillApplyNewRequest 

}

func NewTaobaoWlbWaybillIGetRequest() *TaobaoWlbWaybillIGetRequest{
    return &TaobaoWlbWaybillIGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWaybillIGetRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.get"
}

func (r TaobaoWlbWaybillIGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWaybillIGetRequest) SetWaybillApplyNewRequest(waybillApplyNewRequest *WaybillApplyNewRequest) error {
    r.waybillApplyNewRequest = waybillApplyNewRequest
    r.Set("waybill_apply_new_request", waybillApplyNewRequest)
    return nil
}

func (r TaobaoWlbWaybillIGetRequest) GetWaybillApplyNewRequest() *WaybillApplyNewRequest {
    return r.waybillApplyNewRequest
}

