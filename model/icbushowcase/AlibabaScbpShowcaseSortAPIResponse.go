package icbushowcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseSortAPIResponse 橱窗顺序变更 API返回值
// alibaba.scbp.showcase.sort
//
// 橱窗顺序变更
type AlibabaScbpShowcaseSortAPIResponse struct {
	model.CommonResponse
	AlibabaScbpShowcaseSortAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseSortAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpShowcaseSortAPIResponseModel).Reset()
}

// AlibabaScbpShowcaseSortAPIResponseModel is 橱窗顺序变更 成功返回结果
type AlibabaScbpShowcaseSortAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_sort_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否更新成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseSortAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpShowcaseSortAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpShowcaseSortAPIResponse)
	},
}

// GetAlibabaScbpShowcaseSortAPIResponse 从 sync.Pool 获取 AlibabaScbpShowcaseSortAPIResponse
func GetAlibabaScbpShowcaseSortAPIResponse() *AlibabaScbpShowcaseSortAPIResponse {
	return poolAlibabaScbpShowcaseSortAPIResponse.Get().(*AlibabaScbpShowcaseSortAPIResponse)
}

// ReleaseAlibabaScbpShowcaseSortAPIResponse 将 AlibabaScbpShowcaseSortAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpShowcaseSortAPIResponse(v *AlibabaScbpShowcaseSortAPIResponse) {
	v.Reset()
	poolAlibabaScbpShowcaseSortAPIResponse.Put(v)
}
