package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceDeductGetAPIRequest
发票扣减的接口 API请求
alibaba.einvoice.deduct.get

获取历史发票扣减量、每日发票扣减量的接口 */
type AlibabaEinvoiceDeductGetAPIRequest struct {
	model.Params
	// 税号
	_payeeRegisterNo string
	// 业务日期
	_bizDate string
	// 类型 1：所有 2：当日
	_type int64
}

// New
