package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordStatusUpdateAPIResponse 关键词启动暂停推广 API返回值
// alibaba.scbp.ad.keyword.status.update
//
// 关键词启动暂停推广
type AlibabaScbpAdKeywordStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdKeywordStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdKeywordStatusUpdateAPIResponseModel).Reset()
}

// AlibabaScbpAdKeywordStatusUpdateAPIResponseModel is 关键词启动暂停推广 成功返回结果
type AlibabaScbpAdKeywordStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新关键词推广状态是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdKeywordStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpAdKeywordStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdKeywordStatusUpdateAPIResponse)
	},
}

// GetAlibabaScbpAdKeywordStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaScbpAdKeywordStatusUpdateAPIResponse
func GetAlibabaScbpAdKeywordStatusUpdateAPIResponse() *AlibabaScbpAdKeywordStatusUpdateAPIResponse {
	return poolAlibabaScbpAdKeywordStatusUpdateAPIResponse.Get().(*AlibabaScbpAdKeywordStatusUpdateAPIResponse)
}

// ReleaseAlibabaScbpAdKeywordStatusUpdateAPIResponse 将 AlibabaScbpAdKeywordStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdKeywordStatusUpdateAPIResponse(v *AlibabaScbpAdKeywordStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdKeywordStatusUpdateAPIResponse.Put(v)
}
