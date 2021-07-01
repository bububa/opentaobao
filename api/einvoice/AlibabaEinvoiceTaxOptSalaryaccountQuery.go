package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
查询用户的发薪账号 
alibaba.einvoice.tax.opt.salaryaccount.query

查询用户的发薪账号状态
*/
func AlibabaEinvoiceTaxOptSalaryaccountQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
