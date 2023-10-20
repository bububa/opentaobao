package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagDeleteAPIResponse 删除关键词分组 API返回值
// alibaba.scbp.tag.delete
//
// 删除关键词分组
type AlibabaScbpTagDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTagDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTagDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTagDeleteAPIResponseModel).Reset()
}

// AlibabaScbpTagDeleteAPIResponseModel is 删除关键词分组 成功返回结果
type AlibabaScbpTagDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_tag_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除关键词分组成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTagDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpTagDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTagDeleteAPIResponse)
	},
}

// GetAlibabaScbpTagDeleteAPIResponse 从 sync.Pool 获取 AlibabaScbpTagDeleteAPIResponse
func GetAlibabaScbpTagDeleteAPIResponse() *AlibabaScbpTagDeleteAPIResponse {
	return poolAlibabaScbpTagDeleteAPIResponse.Get().(*AlibabaScbpTagDeleteAPIResponse)
}

// ReleaseAlibabaScbpTagDeleteAPIResponse 将 AlibabaScbpTagDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTagDeleteAPIResponse(v *AlibabaScbpTagDeleteAPIResponse) {
	v.Reset()
	poolAlibabaScbpTagDeleteAPIResponse.Put(v)
}
