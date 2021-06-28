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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_open_isattention_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlibabaInteractOpenIsattentionResultDo `json:"result,omitempty" xml:"