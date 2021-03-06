package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse 查询买赠活动 API返回值
// alibaba.wdk.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel
}

// AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel is 查询买赠活动 成功返回结果
type AlibabaWdkMarketingItembuygiftQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itembuygift_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
