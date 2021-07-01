package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceIncomeScanReturnAPIRequest
进项扫描状态回传 API请求
alibaba.einvoice.income.scan.return

回传进项扫描每个阶段的状态，比如ocr开始，ocr结束，查验开始，查验结束等 */
type AlibabaEinvoiceIncomeScanReturnAPIRequest struct {
	model.Params
	// 扫描的批次号
	_batchNo string
	// 扫描状态，0=开始ocr，1=ocr结束，2=开始查验，3=查验结束
	_status int64
	// 该批次对应的发票数量，扫描结束和查验结束status=1，3时必填
	_invoiceCount int64
	// 驱动是否成功，true=成功，false=失败
	_success bool
	// 错误码，success=false时填入
	_errorCode string
	// 错误信息，success=false时必填
	_errorMessage string
}

// New
