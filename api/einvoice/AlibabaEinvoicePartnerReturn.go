package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
开票商回传开票结果 
alibaba.einvoice.partner.return

开票商返回开票结果数据
*/
func AlibabaEinvoicePartnerReturn(clt *core.SDKClient, req *einvoice.AlibabaEinvoicePartnerReturnRequest, session string) (*einvoice.AlibabaEinvoicePartnerReturnAPIResponse, error) {
    var resp einvoice.AlibabaEinvoicePartnerReturnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
