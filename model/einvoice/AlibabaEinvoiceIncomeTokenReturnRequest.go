package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传税号token APIRequest
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

func NewAlibabaEinvoiceIncomeTokenReturnRequest() *AlibabaEinvoiceIncomeTokenReturnRequest{
    return &AlibabaEinvoiceIncomeTokenReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.token.return"
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetArea() string {
    return r.area
}

func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetErrorCode() string {
    return r.errorCode
}

func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}

func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetExpireTime(expireTime string) error {
    r.expireTime = expireTime
    r.Set("expire_time", expireTime)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetExpireTime() string {
    return r.expireTime
}

func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetPayeeName(payeeName string) error {
    r.payeeName = payeeName
    r.Set("payee_name", payeeName)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetPayeeName() string {
    return r.payeeName
}

func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetSuccess() bool {
    return r.success
}

func (r *AlibabaEinvoiceIncomeTokenReturnRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaEinvoiceIncomeTokenReturnRequest) GetToken() string {
    return r.token
}

