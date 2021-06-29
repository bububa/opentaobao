package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
商家自研ERP开票结果获取 
alibaba.einvoice.merchant.result.get

商家自研ERP开票结果获取
*/
func AlibabaEinvoiceMerchantResultGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantResultGetRequest, session string) (*einvoice.AlibabaEinvoiceMerchantResultGetAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceMerchantResultGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
