package idle

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代扣详情查询 APIResponse
alibaba.idle.agreement.pay.query

查询代扣结果
*/
type AlibabaIdleAgreementPayQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIdleAgreementPayQueryResponse
}

type AlibabaIdleAgreementPayQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_agreement_pay_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回
    
    Result   *AlibabaIdleAgreementPayQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
