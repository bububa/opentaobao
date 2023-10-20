package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedInvalidateAPIResponse aliexpress.solution.feed.invalidate API返回值
// aliexpress.solution.feed.invalidate
//
// Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation.
type AliexpressSolutionFeedInvalidateAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionFeedInvalidateAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedInvalidateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionFeedInvalidateAPIResponseModel).Reset()
}

// AliexpressSolutionFeedInvalidateAPIResponseModel is aliexpress.solution.feed.invalidate 成功返回结果
type AliexpressSolutionFeedInvalidateAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_feed_invalidate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// successful list
	SuccessfulList []int64 `json:"successful_list,omitempty" xml:"successful_list>int64,omitempty"`
	// no need to be invalidated jobs. Usually caused by that the job is being or has already been executed.
	NoNeedInvalidationList []int64 `json:"no_need_invalidation_list,omitempty" xml:"no_need_invalidation_list>int64,omitempty"`
	// failed list
	FailedList []int64 `json:"failed_list,omitempty" xml:"failed_list>int64,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionFeedInvalidateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SuccessfulList = m.SuccessfulList[:0]
	m.NoNeedInvalidationList = m.NoNeedInvalidationList[:0]
	m.FailedList = m.FailedList[:0]
}

var poolAliexpressSolutionFeedInvalidateAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionFeedInvalidateAPIResponse)
	},
}

// GetAliexpressSolutionFeedInvalidateAPIResponse 从 sync.Pool 获取 AliexpressSolutionFeedInvalidateAPIResponse
func GetAliexpressSolutionFeedInvalidateAPIResponse() *AliexpressSolutionFeedInvalidateAPIResponse {
	return poolAliexpressSolutionFeedInvalidateAPIResponse.Get().(*AliexpressSolutionFeedInvalidateAPIResponse)
}

// ReleaseAliexpressSolutionFeedInvalidateAPIResponse 将 AliexpressSolutionFeedInvalidateAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionFeedInvalidateAPIResponse(v *AliexpressSolutionFeedInvalidateAPIResponse) {
	v.Reset()
	poolAliexpressSolutionFeedInvalidateAPIResponse.Put(v)
}
