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
type AlitripBtripInvoiceSettingRuleAPIRequest struct {
    model.Params
    // 入参对象
    _rq   *OpenInvoiceRuleRq
}

// 初始化AlitripBtripInvoiceSettingRuleAPIRequest对象
func NewAlitripBtripInvoiceSettingRuleRequest() *AlitripBtripInvoiceSettingRuleAPIRequest{
    return &AlitripBtripInvoiceSettingRuleAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSettingRuleAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.setting.rule"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSettingRuleAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 入参对象
func (r *AlitripBtripInvoiceSettingRuleAPIRequest) SetRq(_rq *OpenInvoiceRuleRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripInvoiceSettingRuleAPIRequest) GetRq() *OpenInvoiceRuleRq {
    return r._rq
}
