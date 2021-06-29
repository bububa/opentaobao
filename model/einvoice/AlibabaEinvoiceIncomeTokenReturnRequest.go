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
    area   string
    // 错误码，success=false时必填
    errorCode   string
    // 错误信息，success=false时必填
    errorMessage   string
    // token过期时间，success=true时必填
    expireTime   string
    // 销售方企业名称，success=true时必填
    payeeName   string
    // 销售方纳税人识别号，success=true时必填
    payeeRegisterNo   string
    // token是否获取成功，true=成功， false=失败
    success   bool
    // token，success=true时必填
    token   string
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
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

// Area Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetArea() string {
    return r.area
}
// ErrorCode Setter
// 错误码，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

// ErrorCode Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetErrorCode() string {
    return r.errorCode
}
// ErrorMessage Setter
// 错误信息，success=false时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

// ErrorMessage Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}
// ExpireTime Setter
// token过期时间，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetExpireTime(expireTime string) error {
    r.expireTime = expireTime
    r.Set("expire_time", expireTime)
    return nil
}

// ExpireTime Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetExpireTime() string {
    return r.expireTime
}
// PayeeName Setter
// 销售方企业名称，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetPayeeName(payeeName string) error {
    r.payeeName = payeeName
    r.Set("payee_name", payeeName)
    return nil
}

// PayeeName Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetPayeeName() string {
    return r.payeeName
}
// PayeeRegisterNo Setter
// 销售方纳税人识别号，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
// Success Setter
// token是否获取成功，true=成功， false=失败
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

// Success Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetSuccess() bool {
    return r.success
}
// Token Setter
// token，success=true时必填
func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetToken() string {
    return r.token
}
