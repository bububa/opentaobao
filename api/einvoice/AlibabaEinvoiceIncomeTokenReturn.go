package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
服务商回传税号token 
alibaba.einvoice.income.token.return

服务商回传税号token，用来勾选抵扣认证
*/
func AlibabaEinvoiceIncomeTokenReturn(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeTokenReturnRequest, session string) (*einvoice.AlibabaEinvoiceIncomeTokenReturnAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceIncomeTokenReturnAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}