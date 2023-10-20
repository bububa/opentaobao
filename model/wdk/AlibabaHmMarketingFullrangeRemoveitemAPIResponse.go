package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingfullrangeremoveitemAPIResponse 全场活动删除购品 API返回值
// alibaba.hm.marketing.fullrange.removeitem
//
// 删除换购商品
type AlibabahmmarketingfullrangeremoveitemAPIResponse struct {
	model.CommonResponse
	AlibabahmmarketingfullrangeremoveitemAPIResponseModel
}

// AlibabahmmarketingfullrangeremoveitemAPIResponseModel is 全场活动删除购品 成功返回结果
type AlibabahmmarketingfullrangeremoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
