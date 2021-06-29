package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
商家开票申请单状态回传 
alibaba.einvoice.invoiceapply.update

开票服务商更新商家开票申请单状态
*/
func AlibabaEinvoiceInvoiceapplyUpdate(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceInvoiceapplyUpdateRequest, session string) (*einvoice.AlibabaEinvoiceInvoiceapplyUpdateAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceInvoiceapplyUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
