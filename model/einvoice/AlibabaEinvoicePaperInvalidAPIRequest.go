package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoicePaperInvalidAPIRequest
纸票作废接口 API请求
alibaba.einvoice.paper.invalid

作废一张已开具的纸票，开票日期在当月，产生逆向时作废即可，开票日期跨月则冲红蓝票 */
type AlibabaEinvoicePaperInvalidAPIRequest struct {
	model.Params
	// 发票代码，空白作废时必填
	_invoiceCode string
	// 发票号码，空白作废时必填
	_invoiceNo string
	// 作废操作人
	_invalidOperator string
	// 作废类型, 0=空白发票(有残缺 的纸张发票，不能做为有效报销)作废, 1=已开发票作废
	_invalidType int64
	// 销售方纳税人识别号
	_payeeRegisterNo string
	// 开票流水号
	_serialNo string
}

// New
