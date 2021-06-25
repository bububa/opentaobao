package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询面单服务订购及面单使用情况v1.0 APIRequest
taobao.wlb.waybill.i.search

获取发货地&CP开通状态&账户的使用情况
*/
type TaobaoWlbWaybillISearchRequest struct {
    model.Params

    // 查询网点信息
    waybillApplyRequest   *WaybillApplyRequest 

}

func NewTaobaoWlbWaybillISearchRequest() *TaobaoWlbWaybillISearchRequest{
    return &TaobaoWlbWaybillISearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWaybillISearchRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.search"
}

func (r TaobaoWlbWaybillISearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWaybillISearchRequest) SetWaybillApplyRequest(waybillApplyRequest *WaybillApplyRequest) error {
    r.waybillApplyRequest = waybillApplyRequest
    r.Set("waybill_apply_request", waybillApplyRequest)
    return nil
}

func (r TaobaoWlbWaybillISearchRequest) GetWaybillApplyRequest() *WaybillApplyRequest {
    return r.waybillApplyRequest
}

