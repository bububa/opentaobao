package icburfq

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rfq推荐 APIResponse
alibaba.icbu.rfq.recommend

rfq推荐
*/
type AlibabaIcbuRfqRecommendAPIResponse struct {
    model.CommonResponse
    AlibabaIcbuRfqRecommendResponse
}

type AlibabaIcbuRfqRecommendResponse struct {
    XMLName xml.Name `xml:"alibaba_icbu_rfq_recommend_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
