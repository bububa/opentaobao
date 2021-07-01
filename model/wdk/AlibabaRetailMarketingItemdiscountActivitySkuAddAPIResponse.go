package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponse
特价活动新增商品 API返回值
alibaba.retail.marketing.itemdiscount.activity.sku.add

新增或更新活动商品信息【同城零售】 */
type AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponseModel
}

// AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponseModel is 特价活动新增商品 成功返回结果
type AlibabaRetailMarketingItemdiscountActivitySkuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_sku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
