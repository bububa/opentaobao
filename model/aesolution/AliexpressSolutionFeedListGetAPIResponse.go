package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedListGetAPIResponse aliexpress.solution.feed.list.get API返回值
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
type AliexpressSolutionFeedListGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionFeedListGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionFeedListGetAPIResponseModel).Reset()
}

// AliexpressSolutionFeedListGetAPIResponseModel is aliexpress.solution.feed.list.get 成功返回结果
type AliexpressSolutionFeedListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_feed_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	JobList []BatchOperationJobDto `json:"job_list,omitempty" xml:"job_list>batch_operation_job_dto,omitempty"`
	// current page
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// page size
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// total count of jobs submitted for this seller
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// total page based on the total_count and page_size
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.JobList = m.JobList[:0]
	m.CurrentPage = 0
	m.PageSize = 0
	m.TotalCount = 0
	m.TotalPage = 0
}

var poolAliexpressSolutionFeedListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionFeedListGetAPIResponse)
	},
}

// GetAliexpressSolutionFeedListGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionFeedListGetAPIResponse
func GetAliexpressSolutionFeedListGetAPIResponse() *AliexpressSolutionFeedListGetAPIResponse {
	return poolAliexpressSolutionFeedListGetAPIResponse.Get().(*AliexpressSolutionFeedListGetAPIResponse)
}

// ReleaseAliexpressSolutionFeedListGetAPIResponse 将 AliexpressSolutionFeedListGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionFeedListGetAPIResponse(v *AliexpressSolutionFeedListGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionFeedListGetAPIResponse.Put(v)
}
