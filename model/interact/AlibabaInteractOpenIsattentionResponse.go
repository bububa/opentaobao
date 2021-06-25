package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
判断用户是否收藏某个店铺 APIResponse
alibaba.interact.open.isattention

判断用户是否收藏某个店铺
*/
type AlibabaInteractOpenIsattentionAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractOpenIsattentionResponse `json:"alibaba_interact_open_isattention_response,omitempty"`
}

type AlibabaInteractOpenIsattentionResponse struct {

    // result
    Result   *AlibabaInteractOpenIsattentionResultDo `json:"result,omitempty"`

}
