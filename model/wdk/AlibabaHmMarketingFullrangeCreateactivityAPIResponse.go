package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeCreateactivityAPIResponse 创建全场活动 API返回值
// alibaba.hm.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabaHmMarketingFullrangeCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel
}

// AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel is 创建全场活动 成功返回结果
type AlibabaHmMarketingFullrangeCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
