package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingFullrangeCreateactivityAPIResponse
创建全场活动 API返回值
alibaba.wdk.marketing.fullrange.createactivity

创建全场活动 */
type AlibabaWdkMarketingFullrangeCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel
}

// AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel is 创建全场活动 成功返回结果
type AlibabaWdkMarketingFullrangeCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
