package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse 全场增加换购品 API返回值
// alibaba.wdk.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel
}

// AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel is 全场增加换购品 成功返回结果
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_addexchangeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 添加商品返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
