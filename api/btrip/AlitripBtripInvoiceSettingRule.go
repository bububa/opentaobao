package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
发票规则设置 
alitrip.btrip.invoice.setting.rule

发票规则设置
*/
func AlitripBtripInvoiceSettingRule(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingRuleRequest, session string) (*btrip.AlitripBtripInvoiceSettingRuleAPIResponse, error) {
    var resp btrip.AlitripBtripInvoiceSettingRuleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
