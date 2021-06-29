package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
发票中台-同平台取消授权税号适用商户 
alibaba.einvoice.merchant.delete

税号授权给同平台下其他商户使用后，可以使用此接口取消授权，被取消授权的商户失去开票能力
*/
func AlibabaEinvoiceMerchantDelete(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantDeleteRequest, session string) (*einvoice.AlibabaEinvoiceMerchantDeleteAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceMerchantDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
