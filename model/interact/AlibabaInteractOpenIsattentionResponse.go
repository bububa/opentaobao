package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
判断用户是否收藏某个店铺 APIResponse
alibaba.interact.open.isattention

判断用户是否收藏某个店铺
*/
type AlibabaInteractOpenIsattentionAPIResponse struct {
    model.CommonResponse
    AlibabaInteractOpenIsattentionResponse
}

type AlibabaInteractOpenIsattentionResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_open_isattention_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaInteractOpenIsattentionResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
