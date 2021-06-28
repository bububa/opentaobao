package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
差评导入 APIResponse
alibaba.wdk.channel.comment.create

差评导入
*/
type AlibabaWdkChannelCommentCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_channel_comment_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    ApiResult   *AlibabaWdkChannelCommentCreateApiResult `json:"api_result,omitempty" xml:"