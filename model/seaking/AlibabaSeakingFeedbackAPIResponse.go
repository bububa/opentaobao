package seaking

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaSeakingFeedbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSeakingFeedbackAPIResponseModel).Reset()
}

// AlibabaSeakingFeedbackAPIResponseModel is API服务发布成功商品ID回传 成功返回结果
type AlibabaSeakingFeedbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_seaking_feedback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSeakingFeedbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaSeakingFeedbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSeakingFeedbackAPIResponse)
	},
}

// GetAlibabaSeakingFeedbackAPIResponse 从 sync.Pool 获取 AlibabaSeakingFeedbackAPIResponse
func GetAlibabaSeakingFeedbackAPIResponse() *AlibabaSeakingFeedbackAPIResponse {
	return poolAlibabaSeakingFeedbackAPIResponse.Get().(*AlibabaSeakingFeedbackAPIResponse)
}

// ReleaseAlibabaSeakingFeedbackAPIResponse 将 AlibabaSeakingFeedbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSeakingFeedbackAPIResponse(v *AlibabaSeakingFeedbackAPIResponse) {
	v.Reset()
	poolAlibabaSeakingFeedbackAPIResponse.Put(v)
}
