package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeTokenReturnAPIRequest 服务商回传税号token API请求
// alibaba.einvoice.income.token.return
//
// 服务商回传税号token，用来勾选抵扣认证
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

// NewAlibabaEinvoiceIncomeTokenReturnRequest 初始化AlibabaEinvoiceIncomeTokenReturnAPIRequest对象
func NewAlibabaEinvoiceIncomeTokenReturnRequest() *AlibabaEinvoiceIncomeTokenReturnAPIRequest {
	return &AlibabaEinvoiceIncomeTokenReturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.token.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Area Setter
// 税局所在区域，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// Get Area Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetArea() string {
	return r._area
}

// Set is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// Get ErrorCode Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// Set is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// Get ErrorMessage Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// Set is ExpireTime Setter
// token过期时间，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetExpireTime(_expireTime string) error {
	r._expireTime = _expireTime
	r.Set("expire_time", _expireTime)
	return nil
}

// Get ExpireTime Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetExpireTime() string {
	return r._expireTime
}

// Set is PayeeName Setter
// 销售方企业名称，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// Get PayeeName Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// Set is PayeeRegisterNo Setter
// 销售方纳税人识别号，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// Get PayeeRegisterNo Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// Set is Success Setter
// token是否获取成功，true=成功， false=失败
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// Get Success Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetSuccess() bool {
	return r._success
}

// Set is Token Setter
// token，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetToken() string {
	return r._token
}
