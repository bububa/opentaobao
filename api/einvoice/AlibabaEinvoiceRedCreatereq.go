package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
发票冲红接口 
alibaba.einvoice.red.createreq

发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红
*/
func AlibabaEinvoiceRedCreatereq(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceRedCreatereqRequest, session string) (*einvoice.AlibabaEinvoiceRedCreatereqAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceRedCreatereqAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
