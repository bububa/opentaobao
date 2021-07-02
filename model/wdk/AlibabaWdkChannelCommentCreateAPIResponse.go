package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelCommentCreateAPIResponse 差评导入 API返回值
// alibaba.wdk.channel.comment.create
//
// 差评导入
type AlibabaWdkChannelCommentCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelCommentCreateAPIResponseModel
}

// AlibabaWdkChannelCommentCreateAPIResponseModel is 差评导入 成功返回结果
type AlibabaWdkChannelCommentCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_comment_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelCommentCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
