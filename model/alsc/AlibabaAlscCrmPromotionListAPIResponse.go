package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPromotionListAPIResponse 获取促销规则列表 API返回值
// alibaba.alsc.crm.promotion.list
//
// 获取品牌的促销规则列表
type AlibabaAlscCrmPromotionListAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPromotionListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPromotionListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmPromotionListAPIResponseModel).Reset()
}

// AlibabaAlscCrmPromotionListAPIResponseModel is 获取促销规则列表 成功返回结果
type AlibabaAlscCrmPromotionListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_promotion_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmPromotionListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmPromotionListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmPromotionListAPIResponse)
	},
}

// GetAlibabaAlscCrmPromotionListAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmPromotionListAPIResponse
func GetAlibabaAlscCrmPromotionListAPIResponse() *AlibabaAlscCrmPromotionListAPIResponse {
	return poolAlibabaAlscCrmPromotionListAPIResponse.Get().(*AlibabaAlscCrmPromotionListAPIResponse)
}

// ReleaseAlibabaAlscCrmPromotionListAPIResponse 将 AlibabaAlscCrmPromotionListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmPromotionListAPIResponse(v *AlibabaAlscCrmPromotionListAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmPromotionListAPIResponse.Put(v)
}
