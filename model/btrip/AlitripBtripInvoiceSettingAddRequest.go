package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票设置 APIRequest
alitrip.btrip.invoice.setting.add

发票设置
*/
type AlitripBtripInvoiceSettingAddRequest struct {
    model.Params

    // 请求对象
    rq   *OpenInvoiceModifyAndNewRq 

}

func NewAlitripBtripInvoiceSettingAddRequest() *AlitripBtripInvoiceSettingAddRequest{
    return &AlitripBtripInvoiceSettingAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripInvoiceSettingAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.add"
}

func (r AlitripBtripInvoiceSettingAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripInvoiceSettingAddRequest) SetRq(rq *OpenInvoiceModifyAndNewRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripInvoiceSettingAddRequest) GetRq() *OpenInvoiceModifyAndNewRq {
    return r.rq
}

