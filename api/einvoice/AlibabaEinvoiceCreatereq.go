package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
ERP开票请求接口 
alibaba.einvoice.createreq

ERP发起开票请求
*/
func AlibabaEinvoiceCreatereq(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceCreatereqRequest, session string) (*einvoice.AlibabaEinvoiceCreatereqAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceCreatereqAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
