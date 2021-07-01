package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceIncomeTokenReturnAPIRequest
服务商回传税号token API请求
alibaba.einvoice.income.token.return

服务商回传税号token，用来勾选抵扣认证 */
type AlibabaEinvoiceIncomeTokenReturnAPIRequest struct {
	model.Params
	// 税局所在区域，success=true时必填
	_area string
	// 错误码，success=false时必填
	_errorCode string
	// 错误信息，success=false时必填
	_errorMessage string
	// token过期时间，success=true时必填
	_expireTime string
	// 销售方企业名称，success=true时必填
	_payeeName string
	// 销售方纳税人识别号，success=true时必填
	_payeeRegisterNo string
	// token是否获取成功，true=成功， false=失败
	_success bool
	// token，success=true时必填
	_token string
}

// New
