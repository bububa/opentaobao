package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeRemoveitemAPIResponse 全场活动删除购品 API返回值
// alibaba.wdk.marketing.fullrange.removeitem
//
// 删除换购商品
type AlibabaWdkMarketingFullrangeRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel
}

// AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel is 全场活动删除购品 成功返回结果
type AlibabaWdkMarketingFullrangeRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
