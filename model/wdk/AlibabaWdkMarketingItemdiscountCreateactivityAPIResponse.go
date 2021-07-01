package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse
创建商品特价活动 API返回值
alibaba.wdk.marketing.itemdiscount.createactivity

创建商品特价活动 */
type AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItemdiscountCreateactivityAPIResponseModel
}

// AlibabaWdkMarketingItemdiscountCreateactivityAPIResponseModel is 创建商品特价活动 成功返回结果
type AlibabaWdkMarketingItemdiscountCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itemdiscount_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
