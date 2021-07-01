package btrip

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/btrip"
)

/* 
发票变更 
alitrip.btrip.invoice.setting.modify

发票变更
*/
func AlitripBtripInvoiceSettingModify(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSettingModifyAPIRequest, session string) (*btrip.AlitripBtripInvoiceSettingModifyAPIResponse, error) {
    var resp btrip.AlitripBtripInvoiceSettingModifyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
