package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceincometokenreturnAPIRequest 服务商回传税号token API请求
// alibaba.einvoice.income.token.return
//
// 服务商回传税号token，用来勾选抵扣认证
type AlibabaeinvoiceincometokenreturnAPIRequest struct {
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

// NewAlibabaeinvoiceincometokenreturnRequest 初始化AlibabaeinvoiceincometokenreturnAPIRequest对象
func NewAlibabaeinvoiceincometokenreturnRequest() *AlibabaeinvoiceincometokenreturnAPIRequest {
	return &AlibabaeinvoiceincometokenreturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.income.token.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetArea is Area Setter
// 税局所在区域，success=true时必填
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetArea() string {
	return r._area
}

// SetErrorCode is ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetErrorMessage is ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetExpireTime is ExpireTime Setter
// token过期时间，success=true时必填
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetExpireTime(_expireTime string) error {
	r._expireTime = _expireTime
	r.Set("expire_time", _expireTime)
	return nil
}

// GetExpireTime ExpireTime Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetExpireTime() string {
	return r._expireTime
}

// SetPayeeName is PayeeName Setter
// 销售方企业名称，success=true时必填
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetPayeeName(_payeeName string) error {
	r._payeeName = _payeeName
	r.Set("payee_name", _payeeName)
	return nil
}

// GetPayeeName PayeeName Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetPayeeName() string {
	return r._payeeName
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 销售方纳税人识别号，success=true时必填
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}

// SetToken is Token Setter
// token，success=true时必填
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetToken() string {
	return r._token
}

// SetSuccess is Success Setter
// token是否获取成功，true=成功， false=失败
func (r *AlibabaeinvoiceincometokenreturnAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r AlibabaeinvoiceincometokenreturnAPIRequest) GetSuccess() bool {
	return r._success
}
