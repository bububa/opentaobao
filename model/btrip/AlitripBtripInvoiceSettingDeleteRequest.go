package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票删除 APIRequest
alitrip.btrip.invoice.setting.delete

发票删除
*/
type AlitripBtripInvoiceSettingDeleteRequest struct {
    model.Params

    // 入参
    rq   *OpenInvoiceDeleteRq 

}

func NewAlitripBtripInvoiceSettingDeleteRequest() *AlitripBtripInvoiceSettingDeleteRequest{
    return &AlitripBtripInvoiceSettingDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripInvoiceSettingDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.delete"
}

func (r AlitripBtripInvoiceSettingDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripInvoiceSettingDeleteRequest) SetRq(rq *OpenInvoiceDeleteRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripInvoiceSettingDeleteRequest) GetRq() *OpenInvoiceDeleteRq {
    return r.rq
}

