package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
差评导入 APIResponse
alibaba.wdk.channel.comment.create

差评导入
*/
type AlibabaWdkChannelCommentCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkChannelCommentCreateResponse `json:"alibaba_wdk_channel_comment_create_response,omitempty"` 
    AlibabaWdkChannelCommentCreateResponse
}

/* model for simplify = false
type AlibabaWdkChannelCommentCreateResponse struct {

    // 返回结果
    
    ApiResult  *struct {
        AlibabaWdkChannelCommentCreateApiResult  *AlibabaWdkChannelCommentCreateApiResult `json:"alibaba_wdk_channel_comment_create_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkChannelCommentCreateResponse struct {

    // 返回结果
    ApiResult   *AlibabaWdkChannelCommentCreateApiResult `json:"api_result,omitempty"`

}
