package einvoice

import (
	"net/url"
	"sync"

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
	// token，success=true时必填
	_token string
	// token是否获取成功，true=成功， false=失败
	_success bool
}

// NewAlibabaEinvoiceIncomeTokenReturnRequest 初始化AlibabaEinvoiceIncomeTokenReturnAPIRequest对象
func NewAlibabaEinvoiceIncomeTokenReturnRequest() *AlibabaEinvoiceIncomeTokenReturnAPIRequest {
	return &AlibabaEinvoiceIncomeTokenReturnAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) Reset() {
	r._area = ""
	r._errorCode = ""
	r._errorMessage = ""
	r._expireTime = ""
	r._payeeName = ""
	r._payeeRegisterNo = ""
	r._token = ""
	r._success = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.token.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetArea is Area Setter
// 税局所在区域，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetArea() string {
	return r._area
}

// SetErrorCode is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetExpireTime is ExpireTime Setter
// token过期时间，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetExpireTime(_expireTime string) error {
	r._expireTime = _expireTime
	r.Set("expire_time", _expireTime)
	return nil
}

// GetExpireTime ExpireTime Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetExpireTime() string {
	return r._expireTime
}

// SetPayeeName is PayeeName Setter
// 销售方企业名称，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// GetPayeeName PayeeName Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方纳税人识别号，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetToken is Token Setter
// token，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetToken() string {
	return r._token
}

// SetSuccess is Success Setter
// token是否获取成功，true=成功， false=失败
func (r *AlibabaEinvoiceIncomeTokenReturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaEinvoiceIncomeTokenReturnAPIRequest) GetSuccess() bool {
	return r._success
}

var poolAlibabaEinvoiceIncomeTokenReturnAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceIncomeTokenReturnRequest()
	},
}

// GetAlibabaEinvoiceIncomeTokenReturnRequest 从 sync.Pool 获取 AlibabaEinvoiceIncomeTokenReturnAPIRequest
func GetAlibabaEinvoiceIncomeTokenReturnAPIRequest() *AlibabaEinvoiceIncomeTokenReturnAPIRequest {
	return poolAlibabaEinvoiceIncomeTokenReturnAPIRequest.Get().(*AlibabaEinvoiceIncomeTokenReturnAPIRequest)
}

// ReleaseAlibabaEinvoiceIncomeTokenReturnAPIRequest 将 AlibabaEinvoiceIncomeTokenReturnAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceIncomeTokenReturnAPIRequest(v *AlibabaEinvoiceIncomeTokenReturnAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceIncomeTokenReturnAPIRequest.Put(v)
}
