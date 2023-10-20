package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftQueryactivityAPIResponse 查询买赠活动 API返回值
// alibaba.hm.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabaHmMarketingItembuygiftQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel
}

// AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel is 查询买赠活动 成功返回结果
type AlibabaHmMarketingItembuygiftQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itembuygift_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
