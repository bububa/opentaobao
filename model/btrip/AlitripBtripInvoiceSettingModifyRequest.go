package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票变更 APIRequest
alitrip.btrip.invoice.setting.modify

发票变更
*/
type AlitripBtripInvoiceSettingModifyRequest struct {
    model.Params

    // 请求对象
    rq   *OpenInvoiceModifyAndNewRq 

}

func NewAlitripBtripInvoiceSettingModifyRequest() *AlitripBtripInvoiceSettingModifyRequest{
    return &AlitripBtripInvoiceSettingModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripInvoiceSettingModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.modify"
}

func (r AlitripBtripInvoiceSettingModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripInvoiceSettingModifyRequest) SetRq(rq *OpenInvoiceModifyAndNewRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripInvoiceSettingModifyRequest) GetRq() *OpenInvoiceModifyAndNewRq {
    return r.rq
}

