package einvoice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/einvoice"
)

/* 
更新用户发薪资产 
alibaba.einvoice.tax.opt.salaryaccount.update

更新用户的发薪账号
*/
func AlibabaEinvoiceTaxOptSalaryaccountUpdate(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse, error) {
    var resp einvoice.AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
