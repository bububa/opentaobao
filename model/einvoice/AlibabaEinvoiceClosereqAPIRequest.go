package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceClosereqAPIRequest
关闭开票失败请求（失败列表可重试） API请求
alibaba.einvoice.closereq

关闭失败开票请求，避免造成重复开票 */
type AlibabaEinvoiceClosereqAPIRequest struct {
	model.Params
	// 流水号
	_serialNo string
	// 税号
	_payeeRegisterNo string
}

// New
