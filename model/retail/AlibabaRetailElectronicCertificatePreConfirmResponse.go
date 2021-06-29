package retail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机开始核销接口 APIResponse
alibaba.retail.electronic.certificate.pre.confirm

零售终端贩卖机开始核销接口,返回待领的商品ID
*/
type AlibabaRetailElectronicCertificatePreConfirmAPIResponse struct {
    model.CommonResponse
    AlibabaRetailElectronicCertificatePreConfirmResponse
}

type AlibabaRetailElectronicCertificatePreConfirmResponse struct {
    XMLName xml.Name `xml:"alibaba_retail_electronic_certificate_pre_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaRetailElectronicCertificatePreConfirmResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
