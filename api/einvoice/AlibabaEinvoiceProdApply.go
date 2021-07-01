package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
提交发票申请 
alibaba.einvoice.prod.apply

提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
*/
func AlibabaEinvoiceProdApply(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceProdApplyAPIRequest, session string) (*einvoice.AlibabaEinvoiceProdApplyAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceProdApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
