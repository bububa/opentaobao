package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest
查询用户签约税优结果 API请求
alibaba.einvoice.tax.opt.esignresult.query

查询用户是否已经签约 */
type AlibabaEinvoiceTaxOptEsignresultQueryAPIRequest struct {
	model.Params
	// 业务方编码
	_employerCode string
	// 用户在业务方平台的userid
	_identificationInBelongingEmployer string
}

// New
