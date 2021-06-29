package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票规则设置 API请求
alitrip.btrip.invoice.setting.rule

发票规则设置
*/
type AlitripBtripInvoiceSettingRuleRequest struct {
    model.Params
    // 入参对象
    rq   *OpenInvoiceRuleRq
}

// 初始化AlitripBtripInvoiceSettingRuleRequest对象
func NewAlitripBtripInvoiceSettingRuleRequest() *AlitripBtripInvoiceSettingRuleRequest{
    return &AlitripBtripInvoiceSettingRuleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingRuleRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.rule"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingRuleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripInvoiceSettingRuleRequest) SetRq(rq *OpenInvoiceRuleRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripInvoiceSettingRuleRequest) GetRq() *OpenInvoiceRuleRq {
    return r.rq
}
