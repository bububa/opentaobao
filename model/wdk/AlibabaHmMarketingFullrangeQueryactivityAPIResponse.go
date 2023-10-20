package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeQueryactivityAPIResponse 全场活动查询活动 API返回值
// alibaba.hm.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabaHmMarketingFullrangeQueryactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel
}

// AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel is 全场活动查询活动 成功返回结果
type AlibabaHmMarketingFullrangeQueryactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_queryactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
