package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse 删除商品池活动商品【同城零售】 API返回值
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
type AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponseModel is 删除商品池活动商品【同城零售】 成功返回结果
type AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse)
	},
}

// GetAlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse
func GetAlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse() *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse {
	return poolAlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse.Get().(*AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse)
}

// ReleaseAlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse 将 AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse(v *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivitySkuDeleteAPIResponse.Put(v)
}
