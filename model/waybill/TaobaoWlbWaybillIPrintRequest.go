package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
打印确认接口v1.0 APIRequest
taobao.wlb.waybill.i.print

打印面单前的校验接口，判断面单号信息与订单信息是否匹配。
*/
type TaobaoWlbWaybillIPrintRequest struct {
    model.Params

    // 打印请求
    waybillApplyPrintCheckRequest   *WaybillApplyPrintCheckRequest 

}

func NewTaobaoWlbWaybillIPrintRequest() *TaobaoWlbWaybillIPrintRequest{
    return &TaobaoWlbWaybillIPrintRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWaybillIPrintRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.i.print"
}

func (r TaobaoWlbWaybillIPrintRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWaybillIPrintRequest) SetWaybillApplyPrintCheckRequest(waybillApplyPrintCheckRequest *WaybillApplyPrintCheckRequest) error {
    r.waybillApplyPrintCheckRequest = waybillApplyPrintCheckRequest
    r.Set("waybill_apply_print_check_request", waybillApplyPrintCheckRequest)
    return nil
}

func (r TaobaoWlbWaybillIPrintRequest) GetWaybillApplyPrintCheckRequest() *WaybillApplyPrintCheckRequest {
    return r.waybillApplyPrintCheckRequest
}

