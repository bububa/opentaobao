package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse 商品池活动新增商品 API返回值
// alibaba.retail.marketing.itempool.activity.sku.add
//
// 新增或更新商品池活动商品信息【同城零售】
type AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolActivitySkuAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItempoolActivitySkuAddAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItempoolActivitySkuAddAPIResponseModel is 商品池活动新增商品 成功返回结果
type AlibabaRetailMarketingItempoolActivitySkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivitySkuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItempoolActivitySkuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse)
	},
}

// GetAlibabaRetailMarketingItempoolActivitySkuAddAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse
func GetAlibabaRetailMarketingItempoolActivitySkuAddAPIResponse() *AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse {
	return poolAlibabaRetailMarketingItempoolActivitySkuAddAPIResponse.Get().(*AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse)
}

// ReleaseAlibabaRetailMarketingItempoolActivitySkuAddAPIResponse 将 AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivitySkuAddAPIResponse(v *AlibabaRetailMarketingItempoolActivitySkuAddAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivitySkuAddAPIResponse.Put(v)
}
