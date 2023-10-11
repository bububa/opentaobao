package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftCreateactivityAPIResponse 创建买赠活动 API返回值
// alibaba.hm.marketing.itembuygift.createactivity
//
// 创建买赠活动
type AlibabaHmMarketingItembuygiftCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItembuygiftCreateactivityAPIResponseModel
}

// AlibabaHmMarketingItembuygiftCreateactivityAPIResponseModel is 创建买赠活动 成功返回结果
type AlibabaHmMarketingItembuygiftCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
