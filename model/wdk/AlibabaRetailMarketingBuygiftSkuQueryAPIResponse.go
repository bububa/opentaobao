package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftSkuQueryAPIResponse 查询买赠活动商品【同城零售】 API返回值
// alibaba.retail.marketing.buygift.sku.query
//
// 查询买赠活动商品【同城零售】
type AlibabaRetailMarketingBuygiftSkuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftSkuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftSkuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftSkuQueryAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftSkuQueryAPIResponseModel is 查询买赠活动商品【同城零售】 成功返回结果
type AlibabaRetailMarketingBuygiftSkuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_sku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 买赠商品查询结果
	Data []BuyGiftActivitySkuDto `json:"data,omitempty" xml:"data>buy_gift_activity_sku_dto,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrNumber string `json:"err_number,omitempty" xml:"err_number,omitempty"`
	// 分页信息
	PageInfo *PageInfoDto `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// 成功标识
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftSkuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.ErrMessage = ""
	m.ErrNumber = ""
	m.PageInfo = nil
	m.Succeed = false
}

var poolAlibabaRetailMarketingBuygiftSkuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftSkuQueryAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftSkuQueryAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftSkuQueryAPIResponse
func GetAlibabaRetailMarketingBuygiftSkuQueryAPIResponse() *AlibabaRetailMarketingBuygiftSkuQueryAPIResponse {
	return poolAlibabaRetailMarketingBuygiftSkuQueryAPIResponse.Get().(*AlibabaRetailMarketingBuygiftSkuQueryAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftSkuQueryAPIResponse 将 AlibabaRetailMarketingBuygiftSkuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftSkuQueryAPIResponse(v *AlibabaRetailMarketingBuygiftSkuQueryAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftSkuQueryAPIResponse.Put(v)
}
