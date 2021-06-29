package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
服务商订购单上传核对 
alibaba.einvoice.unitorder.check

开票服务商回传收到的订购单用于电子发票平台核对
*/
func AlibabaEinvoiceUnitorderCheck(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceUnitorderCheckRequest, session string) (*einvoice.AlibabaEinvoiceUnitorderCheckAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceUnitorderCheckAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
