package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagListAPIResponse 查询所有分组 API返回值
// alibaba.scbp.tag.list
//
// 查询所有分组
type AlibabaScbpTagListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTagListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTagListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTagListAPIResponseModel).Reset()
}

// AlibabaScbpTagListAPIResponseModel is 查询所有分组 成功返回结果
type AlibabaScbpTagListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_tag_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 所有分组
	TagList []TagGroup `json:"tag_list,omitempty" xml:"tag_list>tag_group,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTagListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TagList = m.TagList[:0]
}

var poolAlibabaScbpTagListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTagListAPIResponse)
	},
}

// GetAlibabaScbpTagListAPIResponse 从 sync.Pool 获取 AlibabaScbpTagListAPIResponse
func GetAlibabaScbpTagListAPIResponse() *AlibabaScbpTagListAPIResponse {
	return poolAlibabaScbpTagListAPIResponse.Get().(*AlibabaScbpTagListAPIResponse)
}

// ReleaseAlibabaScbpTagListAPIResponse 将 AlibabaScbpTagListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTagListAPIResponse(v *AlibabaScbpTagListAPIResponse) {
	v.Reset()
	poolAlibabaScbpTagListAPIResponse.Put(v)
}
