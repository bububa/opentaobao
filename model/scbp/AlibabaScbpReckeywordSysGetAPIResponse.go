package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpReckeywordSysGetAPIResponse 系统推荐 API返回值
// alibaba.scbp.reckeyword.sys.get
//
// 查询系统推荐词
type AlibabaScbpReckeywordSysGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpReckeywordSysGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpReckeywordSysGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpReckeywordSysGetAPIResponseModel).Reset()
}

// AlibabaScbpReckeywordSysGetAPIResponseModel is 系统推荐 成功返回结果
type AlibabaScbpReckeywordSysGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_reckeyword_sys_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统推荐词结果列表
	ResultList []RecKeywordDto `json:"result_list,omitempty" xml:"result_list>rec_keyword_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpReckeywordSysGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultList = m.ResultList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpReckeywordSysGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpReckeywordSysGetAPIResponse)
	},
}

// GetAlibabaScbpReckeywordSysGetAPIResponse 从 sync.Pool 获取 AlibabaScbpReckeywordSysGetAPIResponse
func GetAlibabaScbpReckeywordSysGetAPIResponse() *AlibabaScbpReckeywordSysGetAPIResponse {
	return poolAlibabaScbpReckeywordSysGetAPIResponse.Get().(*AlibabaScbpReckeywordSysGetAPIResponse)
}

// ReleaseAlibabaScbpReckeywordSysGetAPIResponse 将 AlibabaScbpReckeywordSysGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpReckeywordSysGetAPIResponse(v *AlibabaScbpReckeywordSysGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpReckeywordSysGetAPIResponse.Put(v)
}
