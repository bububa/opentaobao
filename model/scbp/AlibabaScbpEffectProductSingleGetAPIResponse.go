package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpEffectProductSingleGetAPIResponse 单个产品的报表 API返回值
// alibaba.scbp.effect.product.single.get
//
// 单个产品的报表
type AlibabaScbpEffectProductSingleGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpEffectProductSingleGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpEffectProductSingleGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpEffectProductSingleGetAPIResponseModel).Reset()
}

// AlibabaScbpEffectProductSingleGetAPIResponseModel is 单个产品的报表 成功返回结果
type AlibabaScbpEffectProductSingleGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_effect_product_single_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单个产品的效果数据列表
	SProductEffectList []SingleProductEffectDto `json:"s_product_effect_list,omitempty" xml:"s_product_effect_list>single_product_effect_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpEffectProductSingleGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SProductEffectList = m.SProductEffectList[:0]
	m.TotalNum = 0
	m.TotalPage = 0
}

var poolAlibabaScbpEffectProductSingleGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpEffectProductSingleGetAPIResponse)
	},
}

// GetAlibabaScbpEffectProductSingleGetAPIResponse 从 sync.Pool 获取 AlibabaScbpEffectProductSingleGetAPIResponse
func GetAlibabaScbpEffectProductSingleGetAPIResponse() *AlibabaScbpEffectProductSingleGetAPIResponse {
	return poolAlibabaScbpEffectProductSingleGetAPIResponse.Get().(*AlibabaScbpEffectProductSingleGetAPIResponse)
}

// ReleaseAlibabaScbpEffectProductSingleGetAPIResponse 将 AlibabaScbpEffectProductSingleGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpEffectProductSingleGetAPIResponse(v *AlibabaScbpEffectProductSingleGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpEffectProductSingleGetAPIResponse.Put(v)
}
