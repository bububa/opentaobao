package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoicePartnerUploadAPIRequest
服务商发票上传接口（非授权） API请求
alibaba.einvoice.partner.upload

服务商发票上传接口（非授权） */
type AlibabaEinvoicePartnerUploadAPIRequest struct {
	model.Params
	// 原蓝票发票号码
	_normalInvoiceNo string
	// 原蓝票发票代码
	_normalInvoiceCode string
	// 销方税号
	_payeeRegisterNo string
	// 发票数据，upload_type=0且invoiceKind=0电子发票时必填
	_invoiceFileData *model.File
	// 发票号码，upload_type=0时必填
	_invoiceNo string
	// 发票代码，upload_type=0时必填
	_invoiceCode string
	// 开票日期，upload_type=0时必填
	_invoiceDate string
	// 密码区
	_cipherText string
	// 机器编号
	_deviceNo string
	// 校验码
	_antiFakeCode string
	// 发票类型，upload_type=0且invoiceKind=0电子发票时必填，暂时只支持pdf
	_fileDataType string
	// 原蓝票的reqIndex
	_reqIndex string
	// 发票种类，0=电子发票，1=纸质普票，2=纸质专票
	_invoiceKind int64
	// 上传的类型，0=冲红上传，1=作废上传
	_uploadType int64
}

// New
