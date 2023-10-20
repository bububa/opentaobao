package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedSubmitAPIResponse aliexpress.solution.feed.submit API返回值
// aliexpress.solution.feed.submit
//
// API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed.
type AliexpressSolutionFeedSubmitAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionFeedSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionFeedSubmitAPIResponseModel).Reset()
}

// AliexpressSolutionFeedSubmitAPIResponseModel is aliexpress.solution.feed.submit 成功返回结果
type AliexpressSolutionFeedSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_feed_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// job id,which is for querying the job response later.
	JobId int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.JobId = 0
}

var poolAliexpressSolutionFeedSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionFeedSubmitAPIResponse)
	},
}

// GetAliexpressSolutionFeedSubmitAPIResponse 从 sync.Pool 获取 AliexpressSolutionFeedSubmitAPIResponse
func GetAliexpressSolutionFeedSubmitAPIResponse() *AliexpressSolutionFeedSubmitAPIResponse {
	return poolAliexpressSolutionFeedSubmitAPIResponse.Get().(*AliexpressSolutionFeedSubmitAPIResponse)
}

// ReleaseAliexpressSolutionFeedSubmitAPIResponse 将 AliexpressSolutionFeedSubmitAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionFeedSubmitAPIResponse(v *AliexpressSolutionFeedSubmitAPIResponse) {
	v.Reset()
	poolAliexpressSolutionFeedSubmitAPIResponse.Put(v)
}
