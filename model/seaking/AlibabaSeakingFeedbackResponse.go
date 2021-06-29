package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
API服务发布成功商品ID回传 APIResponse
alibaba.seaking.feedback

API服务发布成功商品ID回传，用于跟进商品id后续的使用情况
*/
type AlibabaSeakingFeedbackAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingFeedbackResponse
}

type AlibabaSeakingFeedbackResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_feedback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
