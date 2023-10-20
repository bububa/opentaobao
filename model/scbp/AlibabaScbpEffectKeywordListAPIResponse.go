package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectKeywordListAPIResponse 关键词报表 API返回值
// alibaba.scbp.effect.keyword.list
//
// 关键词报表
type AlibabaScbpEffectKeywordListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpEffectKeywordListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpEffectKeywordListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpEffectKeywordListAPIResponseModel).Reset()
}

// AlibabaScbpEffectKeywordListAPIResponseModel is 关键词报表 成功返回结果
type AlibabaScbpEffectKeywordListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_effect_keyword_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词效果数据列表
	KeywordReportList []AdKeywordEffectDto `json:"keyword_report_list,omitempty" xml:"keyword_report_list>ad_keyword_effect_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpEffectKeywordListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.KeywordReportList = m.KeywordReportList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpEffectKeywordListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpEffectKeywordListAPIResponse)
	},
}

// GetAlibabaScbpEffectKeywordListAPIResponse 从 sync.Pool 获取 AlibabaScbpEffectKeywordListAPIResponse
func GetAlibabaScbpEffectKeywordListAPIResponse() *AlibabaScbpEffectKeywordListAPIResponse {
	return poolAlibabaScbpEffectKeywordListAPIResponse.Get().(*AlibabaScbpEffectKeywordListAPIResponse)
}

// ReleaseAlibabaScbpEffectKeywordListAPIResponse 将 AlibabaScbpEffectKeywordListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpEffectKeywordListAPIResponse(v *AlibabaScbpEffectKeywordListAPIResponse) {
	v.Reset()
	poolAlibabaScbpEffectKeywordListAPIResponse.Put(v)
}
