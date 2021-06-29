package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票规则设置 APIRequest
alitrip.btrip.invoice.setting.rule

发票规则设置
*/
type AlitripBtripInvoiceSettingRuleRequest struct {
    model.Params

    // 入参对象
    rq   *OpenInvoiceRuleRq 

}

func NewAlitripBtripInvoiceSettingRuleRequest() *AlitripBtripInvoiceSettingRuleRequest{
    return &AlitripBtripInvoiceSettingRuleRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripInvoiceSettingRuleRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.rule"
}

func (r AlitripBtripInvoiceSettingRuleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripInvoiceSettingRuleRequest) SetRq(rq *OpenInvoiceRuleRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripInvoiceSettingRuleRequest) GetRq() *OpenInvoiceRuleRq {
    return r.rq
}

