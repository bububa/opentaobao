package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传税号token API请求
alibaba.einvoice.income.token.return

服务商回传税号token，用来勾选抵扣认证
*/
type AlibabaEinvoiceIncomeTokenReturnRequest struct {
    model.Params
    // 税局所在区域，success=true时必填
    _area   string
    // 错误码，success=false时必填
    _errorCode   string
    // 错误信息，success=false时必填
    _errorMessage   string
    // token过期时间，success=true时必填
    _expireTime   string
    // 销售方企业名称，success=true时必填
    _payeeName   string
    // 销售方纳税人识别号，success=true时必填
    _payeeRegisterNo   string
    // token是否获取成功，true=成功， false=失败
    _success   bool
    // token，success=true时必填
    _token   string
}

// 初始化AlibabaEinvoiceIncomeTokenReturnRequest对象
func NewAlibabaEinvoiceIncomeTokenReturnRequest() *AlibabaEinvoiceIncomeTokenReturnRequest{
    return &AlibabaEinvoiceIncomeTokenReturnRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.token.return"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Area Setter
// 税局所在区域，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetArea(_area string) error {
    r._area = _area
    r.Set("area", _area)
    return nil
}

// Area Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetArea() string {
    return r._area
}
// ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetErrorCode(_errorCode string) error {
    r._errorCode = _errorCode
    r.Set("error_code", _errorCode)
    return nil
}

// ErrorCode Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetErrorCode() string {
    return r._errorCode
}
// ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetErrorMessage(_errorMessage string) error {
    r._errorMessage = _errorMessage
    r.Set("error_message", _errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetErrorMessage() string {
    return r._errorMessage
}
// ExpireTime Setter
// token过期时间，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetExpireTime(_expireTime string) error {
    r._expireTime = _expireTime
    r.Set("expire_time", _expireTime)
    return nil
}

// ExpireTime Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetExpireTime() string {
    return r._expireTime
}
// PayeeName Setter
// 销售方企业名称，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetPayeeName(_payeeName string) error {
    r._payeeName = _payeeName
    r.Set("payee_name", _payeeName)
    return nil
}

// PayeeName Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetPayeeName() string {
    return r._payeeName
}
// PayeeRegisterNo Setter
// 销售方纳税人识别号，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// Success Setter
// token是否获取成功，true=成功， false=失败
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetSuccess(_success bool) error {
    r._success = _success
    r.Set("success", _success)
    return nil
}

// Success Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetSuccess() bool {
    return r._success
}
// Token Setter
// token，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetToken() string {
    return r._token
}
