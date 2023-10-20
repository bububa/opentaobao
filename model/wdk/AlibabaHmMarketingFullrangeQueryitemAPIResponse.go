package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeQueryitemAPIResponse 全场活动查询换购品 API返回值
// alibaba.hm.marketing.fullrange.queryitem
//
// 全场活动查询换购品
type AlibabaHmMarketingFullrangeQueryitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeQueryitemAPIResponseModel
}

// AlibabaHmMarketingFullrangeQueryitemAPIResponseModel is 全场活动查询换购品 成功返回结果
type AlibabaHmMarketingFullrangeQueryitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_queryitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
