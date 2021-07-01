package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceRedCreatereqAPIRequest
发票冲红接口 API请求
alibaba.einvoice.red.createreq

发票冲红接口，通过蓝票流水号或者发票号码+发票代码进行冲红 */
type AlibabaEinvoiceRedCreatereqAPIRequest struct {
	model.Params
	// 销售方税号
	_payeeRegisterNo string
	// 蓝票流水号，优先级高于发票代码+发票号码
	_blueSerialNo string
	// 红票流水号
	_redSerialNo string
	// 蓝票发票代码
	_invoiceCode string
	// 蓝票发票号码
	_invoiceNo string
}

// New
