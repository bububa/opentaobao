package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceAmountCheckAPIRequest
开票量核对接口 API请求
alibaba.einvoice.amount.check

跟开票服务商核对历史开票量，用来对账 */
type AlibabaEinvoiceAmountCheckAPIRequest struct {
	model.Params
	// 税号
	_payeeRegisterNo string
	// 开票日期开始时间
	_startDate string
	// 开票日期结束时间
	_endDate string
}

// New
