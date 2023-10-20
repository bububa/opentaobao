package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse 查询单品特价活动【同城零售】 API返回值
// alibaba.retail.marketing.itemdiscount.activity.query
//
// 查询单品特价活动【同城零售】
type AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItemdiscountActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItemdiscountActivityQueryAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItemdiscountActivityQueryAPIResponseModel is 查询单品特价活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItemdiscountActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrNumber string `json:"err_number,omitempty" xml:"err_number,omitempty"`
	// 响应体
	Data *ItemDiscountPromotionActivityDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功标识
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItemdiscountActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrMessage = ""
	m.ErrNumber = ""
	m.Data = nil
	m.Succeed = false
}

var poolAlibabaRetailMarketingItemdiscountActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse)
	},
}

// GetAlibabaRetailMarketingItemdiscountActivityQueryAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse
func GetAlibabaRetailMarketingItemdiscountActivityQueryAPIResponse() *AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse {
	return poolAlibabaRetailMarketingItemdiscountActivityQueryAPIResponse.Get().(*AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse)
}

// ReleaseAlibabaRetailMarketingItemdiscountActivityQueryAPIResponse 将 AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItemdiscountActivityQueryAPIResponse(v *AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItemdiscountActivityQueryAPIResponse.Put(v)
}
