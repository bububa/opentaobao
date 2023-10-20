package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptsalaryaccountupdate 更新用户发薪资产
// alibaba.einvoice.tax.opt.salaryaccount.update
//
// 更新用户的发薪账号
func Alibabaeinvoicetaxoptsalaryaccountupdate(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptsalaryaccountupdateAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptsalaryaccountupdateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptsalaryaccountupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
