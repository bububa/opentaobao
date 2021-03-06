package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse 全场活动删除活动接口 API返回值
// alibaba.wdk.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
type AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel
}

// AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel is 全场活动删除活动接口 成功返回结果
type AlibabaWdkMarketingFullrangeDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
