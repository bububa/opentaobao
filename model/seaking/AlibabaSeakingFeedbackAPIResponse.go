package seaking

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingFeedbackAPIResponse API服务发布成功商品ID回传 API返回值
// alibaba.seaking.feedback
//
// API服务发布成功商品ID回传，用于跟进商品id后续的使用情况
type AlibabaSeakingFeedbackAPIResponse struct {
	model.CommonResponse
	AlibabaSeakingFeedbackAPIResponseModel
}

// AlibabaSeakingFeedbackAPIResponseModel is API服务发布成功商品ID回传 成功返回结果
type AlibabaSeakingFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
