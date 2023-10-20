package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountCreateactivityAPIResponse 创建商品特价活动 API返回值
// alibaba.hm.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
type AlibabaHmMarketingItemdiscountCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItemdiscountCreateactivityAPIResponseModel
}

// AlibabaHmMarketingItemdiscountCreateactivityAPIResponseModel is 创建商品特价活动 成功返回结果
type AlibabaHmMarketingItemdiscountCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itemdiscount_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
