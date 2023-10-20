package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeRemoveitemAPIResponse 全场活动删除购品 API返回值
// alibaba.hm.marketing.fullrange.removeitem
//
// 删除换购商品
type AlibabaHmMarketingFullrangeRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel
}

// AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel is 全场活动删除购品 成功返回结果
type AlibabaHmMarketingFullrangeRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
