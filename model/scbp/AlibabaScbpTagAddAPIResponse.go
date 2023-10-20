package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagAddAPIResponse 创建关键词分组 API返回值
// alibaba.scbp.tag.add
//
// 创建关键词分组
type AlibabaScbpTagAddAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTagAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTagAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTagAddAPIResponseModel).Reset()
}

// AlibabaScbpTagAddAPIResponseModel is 创建关键词分组 成功返回结果
type AlibabaScbpTagAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_tag_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分组名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTagAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TagName = ""
}

var poolAlibabaScbpTagAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTagAddAPIResponse)
	},
}

// GetAlibabaScbpTagAddAPIResponse 从 sync.Pool 获取 AlibabaScbpTagAddAPIResponse
func GetAlibabaScbpTagAddAPIResponse() *AlibabaScbpTagAddAPIResponse {
	return poolAlibabaScbpTagAddAPIResponse.Get().(*AlibabaScbpTagAddAPIResponse)
}

// ReleaseAlibabaScbpTagAddAPIResponse 将 AlibabaScbpTagAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTagAddAPIResponse(v *AlibabaScbpTagAddAPIResponse) {
	v.Reset()
	poolAlibabaScbpTagAddAPIResponse.Put(v)
}
