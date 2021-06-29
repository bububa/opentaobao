package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼平台商户代扣 APIResponse
alibaba.idle.agreement.pay

闲鱼平台代扣能力：用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
*/
type AlibabaIdleAgreementPayAPIResponse struct {
    model.CommonResponse
    AlibabaIdleAgreementPayResponse
}

type AlibabaIdleAgreementPayResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_agreement_pay_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaIdleAgreementPayResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
