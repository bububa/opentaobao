package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
发票设置 
alitrip.btrip.invoice.setting.add

发票设置
*/
func AlitripBtripInvoiceSettingAdd(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingAddRequest, session string) (*btrip.AlitripBtripInvoiceSettingAddAPIResponse, error) {
    var resp btrip.AlitripBtripInvoiceSettingAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
