package icbushowcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseListAPIResponse 橱窗查询 API返回值
// alibaba.scbp.showcase.list
//
// 橱窗查询
type AlibabaScbpShowcaseListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpShowcaseListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpShowcaseListAPIResponseModel).Reset()
}

// AlibabaScbpShowcaseListAPIResponseModel is 橱窗查询 成功返回结果
type AlibabaScbpShowcaseListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []Showcase `json:"results,omitempty" xml:"results>showcase,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolAlibabaScbpShowcaseListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpShowcaseListAPIResponse)
	},
}

// GetAlibabaScbpShowcaseListAPIResponse 从 sync.Pool 获取 AlibabaScbpShowcaseListAPIResponse
func GetAlibabaScbpShowcaseListAPIResponse() *AlibabaScbpShowcaseListAPIResponse {
	return poolAlibabaScbpShowcaseListAPIResponse.Get().(*AlibabaScbpShowcaseListAPIResponse)
}

// ReleaseAlibabaScbpShowcaseListAPIResponse 将 AlibabaScbpShowcaseListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpShowcaseListAPIResponse(v *AlibabaScbpShowcaseListAPIResponse) {
	v.Reset()
	poolAlibabaScbpShowcaseListAPIResponse.Put(v)
}
