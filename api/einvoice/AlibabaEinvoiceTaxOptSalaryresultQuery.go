package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptsalaryresultquery 查询发薪结果
// alibaba.einvoice.tax.opt.salaryresult.query
//
// 查询发薪结果
func Alibabaeinvoicetaxoptsalaryresultquery(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptsalaryresultqueryAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptsalaryresultqueryAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptsalaryresultqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
