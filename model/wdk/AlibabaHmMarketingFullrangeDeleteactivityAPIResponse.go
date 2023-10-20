package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeDeleteactivityAPIResponse 全场活动删除活动接口 API返回值
// alibaba.hm.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
type AlibabaHmMarketingFullrangeDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingFullrangeDeleteactivityAPIResponseModel
}

// AlibabaHmMarketingFullrangeDeleteactivityAPIResponseModel is 全场活动删除活动接口 成功返回结果
type AlibabaHmMarketingFullrangeDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_fullrange_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
