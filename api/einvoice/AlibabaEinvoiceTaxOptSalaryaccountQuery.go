package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptsalaryaccountquery 查询用户的发薪账号
// alibaba.einvoice.tax.opt.salaryaccount.query
//
// 查询用户的发薪账号状态
func Alibabaeinvoicetaxoptsalaryaccountquery(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptsalaryaccountqueryAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptsalaryaccountqueryAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptsalaryaccountqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
