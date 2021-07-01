package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest
查询发薪结果 API请求
alibaba.einvoice.tax.opt.salaryresult.query

查询发薪结果 */
type AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest struct {
	model.Params
	// 发薪流水号
	_detailIdList []string
	// 业务方编码
	_employerCode string
}

// New
