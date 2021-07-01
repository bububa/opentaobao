package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
纸质发票结果回传 
alibaba.einvoice.paper.return

纸质发票结果回传
*/
func AlibabaEinvoicePaperReturn(clt *core.SDKClient, req *einvoice.AlibabaEinvoicePaperReturnAPIRequest, session string) (*einvoice.AlibabaEinvoicePaperReturnAPIResponse, error) {
    var resp einvoice.AlibabaEinvoicePaperReturnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
