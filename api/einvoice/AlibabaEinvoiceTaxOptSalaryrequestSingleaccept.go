package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
单明细发薪受理 
alibaba.einvoice.tax.opt.salaryrequest.singleaccept

单明细发薪受理
*/
func AlibabaEinvoiceTaxOptSalaryrequestSingleaccept(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
