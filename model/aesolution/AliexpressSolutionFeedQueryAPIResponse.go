package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedQueryAPIResponse aliexpress.solution.feed.query API返回值
// aliexpress.solution.feed.query
//
// API for query the execution result of feed.
type AliexpressSolutionFeedQueryAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionFeedQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionFeedQueryAPIResponseModel).Reset()
}

// AliexpressSolutionFeedQueryAPIResponseModel is aliexpress.solution.feed.query 成功返回结果
type AliexpressSolutionFeedQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_feed_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Result list after all the item_content,which were previously submitted through API:aliexpress.solution.feed.submit, have been executed , including both successful and unsuccessful items.
	ResultList []SingleItemResponseDto `json:"result_list,omitempty" xml:"result_list>single_item_response_dto,omitempty"`
	// job id
	JobId int64 `json:"job_id,omitempty" xml:"job_id,omitempty"`
	// Count of successful items after executed under this job
	SuccessItemCount int64 `json:"success_item_count,omitempty" xml:"success_item_count,omitempty"`
	// Count of total items under this job
	TotalItemCount int64 `json:"total_item_count,omitempty" xml:"total_item_count,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.JobId = 0
	m.SuccessItemCount = 0
	m.TotalItemCount = 0
}

var poolAliexpressSolutionFeedQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionFeedQueryAPIResponse)
	},
}

// GetAliexpressSolutionFeedQueryAPIResponse 从 sync.Pool 获取 AliexpressSolutionFeedQueryAPIResponse
func GetAliexpressSolutionFeedQueryAPIResponse() *AliexpressSolutionFeedQueryAPIResponse {
	return poolAliexpressSolutionFeedQueryAPIResponse.Get().(*AliexpressSolutionFeedQueryAPIResponse)
}

// ReleaseAliexpressSolutionFeedQueryAPIResponse 将 AliexpressSolutionFeedQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionFeedQueryAPIResponse(v *AliexpressSolutionFeedQueryAPIResponse) {
	v.Reset()
	poolAliexpressSolutionFeedQueryAPIResponse.Put(v)
}
