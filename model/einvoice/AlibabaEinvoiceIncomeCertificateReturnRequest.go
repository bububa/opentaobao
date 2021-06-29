package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商回传进项认证结果 APIRequest
alibaba.einvoice.income.certificate.return

服务商回传客户端agent所处环境的设备列表，比如扫描仪
*/
type AlibabaEinvoiceIncomeCertificateReturnRequest struct {
    model.Params

    // 错误码，success=false时必填
    errorCode   string 

    // 错误信息，success=false时必填
    errorMessage   string 

    // 请求标识
    reqIndex   string 

    // 认证结果，true=成功，false=失败
    success   bool 

    // 认证步骤，1=勾选，2=汇总，3=确认
    step   int64 

}

func NewAlibabaEinvoiceIncomeCertificateReturnRequest() *AlibabaEinvoiceIncomeCertificateReturnRequest{
    return &AlibabaEinvoiceIncomeCertificateReturnRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceIncomeCertificateReturnRequest) GetApiMethodName() string {
    return "alibaba.einvoice.income.certificate.return"
}

func (r AlibabaEinvoiceIncomeCertificateReturnRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceIncomeCertificateReturnRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

func (r AlibabaEinvoiceIncomeCertificateReturnRequest) GetErrorCode() string {
    return r.errorCode
}

func (r *AlibabaEinvoiceIncomeCertificateReturnRequest) SetErrorMessage(errorMessage string) error {
    r.errorMessage = errorMessage
    r.Set("error_message", errorMessage)
    return nil
}

func (r AlibabaEinvoiceIncomeCertificateReturnRequest) GetErrorMessage() string {
    return r.errorMessage
}

func (r *AlibabaEinvoiceIncomeCertificateReturnRequest) SetReqIndex(reqIndex string) error {
    r.reqIndex = reqIndex
    r.Set("req_index", reqIndex)
    return nil
}

func (r AlibabaEinvoiceIncomeCertificateReturnRequest) GetReqIndex() string {
    return r.reqIndex
}

func (r *AlibabaEinvoiceIncomeCertificateReturnRequest) SetSuccess(success bool) error {
    r.success = success
    r.Set("success", success)
    return nil
}

func (r AlibabaEinvoiceIncomeCertificateReturnRequest) GetSuccess() bool {
    return r.success
}

func (r *AlibabaEinvoiceIncomeCertificateReturnRequest) SetStep(step int64) error {
    r.step = step
    r.Set("step", step)
    return nil
}

func (r AlibabaEinvoiceIncomeCertificateReturnRequest) GetStep() int64 {
    return r.step
}

