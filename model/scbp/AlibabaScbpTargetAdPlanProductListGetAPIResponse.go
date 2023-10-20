package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanProductListGetAPIResponse 定向推广-获取推广计划产品列表 API返回值
// alibaba.scbp.target.ad.plan.product.list.get
//
// 定向推广-获取推广计划产品列表
type AlibabaScbpTargetAdPlanProductListGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanProductListGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanProductListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanProductListGetAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanProductListGetAPIResponseModel is 定向推广-获取推广计划产品列表 成功返回结果
type AlibabaScbpTargetAdPlanProductListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_product_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// TopP4pQuickCampaignProductView
	ProductList []TopP4pQuickCampaignProductView `json:"product_list,omitempty" xml:"product_list>top_p4p_quick_campaign_product_view,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanProductListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductList = m.ProductList[:0]
	m.TotalPage = 0
	m.TotalNum = 0
}

var poolAlibabaScbpTargetAdPlanProductListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanProductListGetAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanProductListGetAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanProductListGetAPIResponse
func GetAlibabaScbpTargetAdPlanProductListGetAPIResponse() *AlibabaScbpTargetAdPlanProductListGetAPIResponse {
	return poolAlibabaScbpTargetAdPlanProductListGetAPIResponse.Get().(*AlibabaScbpTargetAdPlanProductListGetAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanProductListGetAPIResponse 将 AlibabaScbpTargetAdPlanProductListGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanProductListGetAPIResponse(v *AlibabaScbpTargetAdPlanProductListGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanProductListGetAPIResponse.Put(v)
}
