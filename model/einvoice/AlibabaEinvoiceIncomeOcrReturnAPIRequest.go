package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceIncomeOcrReturnAPIRequest
服务商回传发票ocr的结果 API请求
alibaba.einvoice.income.ocr.return

服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传 */
type AlibabaEinvoiceIncomeOcrReturnAPIRequest struct {
	model.Params
	// 校验码，ocr结果为普票，success=true并且invoiceKind=1时必填
	_checksum string
	// 错误码，success=false是必填
	_errorCode string
	// 错误消息，success=false是必填
	_errorMessage string
	// 发票ocr影像文件，type=1时必填
	_imageData *model.File
	// 发票ocr影像编号，type=1时必填
	_imageId string
	// 发票代码，success=true时必填
	_invoiceCode string
	// 开票日期，格式为yyyy-MM-dd，success=true时必填
	_invoiceDate string
	// 发票种类，1=普票，2=专票，success=true时必填
	_invoiceKind int64
	// 发票号码，success=true时必填
	_invoiceNo string
	// 开票请求标识，扫描驱动回传type=1时填批次号
	_reqIndex string
	// ocr结果，true=成功，false=失败
	_success bool
	// 不含税金额，ocr结果为专票，success=true并且invoiceKind=2时必填
	_sumPrice string
	// 请求类型，0=阿里主动发起的cor，1=扫描驱动服务商主动回传ocr结果
	_type int64
}

// New
