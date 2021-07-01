package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceIncomeCertificateReturnAPIRequest
服务商回传进项认证结果 API请求
alibaba.einvoice.income.certificate.return

服务商回传客户端agent所处环境的设备列表，比如扫描仪 */
type AlibabaEinvoiceIncomeCertificateReturnAPIRequest struct {
	model.Params
	// 错误码，success=false时必填
	_errorCode string
	// 错误信息，success=false时必填
	_errorMessage string
	// 请求标识
	_reqIndex string
	// 认证结果，true=成功，false=失败
	_success bool
	// 认证步骤，1=勾选，2=汇总，3=确认
	_step int64
}

// New
