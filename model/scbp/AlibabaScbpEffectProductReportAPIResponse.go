package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectProductReportAPIResponse 所有产品报表 API返回值
// alibaba.scbp.effect.product.report
//
// 所有产品报表
type AlibabaScbpEffectProductReportAPIResponse struct {
	model.CommonResponse
	AlibabaScbpEffectProductReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpEffectProductReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpEffectProductReportAPIResponseModel).Reset()
}

// AlibabaScbpEffectProductReportAPIResponseModel is 所有产品报表 成功返回结果
type AlibabaScbpEffectProductReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_effect_product_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品效果数据列表
	ProductEffectList []ProductEffectDto `json:"product_effect_list,omitempty" xml:"product_effect_list>product_effect_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpEffectProductReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductEffectList = m.ProductEffectList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpEffectProductReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpEffectProductReportAPIResponse)
	},
}

// GetAlibabaScbpEffectProductReportAPIResponse 从 sync.Pool 获取 AlibabaScbpEffectProductReportAPIResponse
func GetAlibabaScbpEffectProductReportAPIResponse() *AlibabaScbpEffectProductReportAPIResponse {
	return poolAlibabaScbpEffectProductReportAPIResponse.Get().(*AlibabaScbpEffectProductReportAPIResponse)
}

// ReleaseAlibabaScbpEffectProductReportAPIResponse 将 AlibabaScbpEffectProductReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpEffectProductReportAPIResponse(v *AlibabaScbpEffectProductReportAPIResponse) {
	v.Reset()
	poolAlibabaScbpEffectProductReportAPIResponse.Put(v)
}
