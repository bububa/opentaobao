package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查面单号状态v1.0 APIRequest
taobao.wlb.waybill.i.querydetail

查看面单号的当前状态，如签收、发货、失效等。
*/
type TaobaoWlbWaybillIQuerydetailRequest struct {
    model.Params

    // 面单查询请求
    waybillDetailQueryRequest   *WaybillDetailQueryRequest 

}

func NewTaobaoWlbWaybillIQuerydetailRequest() *TaobaoWlbWaybillIQuerydetailRequest{
    return &TaobaoWlbWaybillIQuerydetailRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWaybillIQuerydetailRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.querydetail"
}

func (r TaobaoWlbWaybillIQuerydetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWaybillIQuerydetailRequest) SetWaybillDetailQueryRequest(waybillDetailQueryRequest *WaybillDetailQueryRequest) error {
    r.waybillDetailQueryRequest = waybillDetailQueryRequest
    r.Set("waybill_detail_query_request", waybillDetailQueryRequest)
    return nil
}

func (r TaobaoWlbWaybillIQuerydetailRequest) GetWaybillDetailQueryRequest() *WaybillDetailQueryRequest {
    return r.waybillDetailQueryRequest
}

