package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoicePaperReturnAPIRequest
纸质发票结果回传 API请求
alibaba.einvoice.paper.return

纸质发票结果回传 */
type AlibabaEinvoicePaperReturnAPIRequest struct {
	model.Params
	// 发票密文，密码区的字符串
	_ciphertext string
	// 发票号码
	_invoiceNo string
	// 发票日期
	_invoiceDate string
	// 防伪码
	_antiFakeCode string
	// 税控设备编号(新版电子发票有)
	_deviceNo string
	// 发票代码
	_invoiceCode string
	// 开票结果"success"或者"fail"
	_createResult string
	// 错误码
	_bizErrorCode string
	// 错误信息
	_bizErrorMsg string
	// 开票请求的唯一索引
	_reqIndex string
}

// New
