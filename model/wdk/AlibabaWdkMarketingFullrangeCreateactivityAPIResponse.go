package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingfullrangecreateactivityAPIResponse 创建全场活动 API返回值
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabawdkmarketingfullrangecreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingfullrangecreateactivityAPIResponseModel
}

// AlibabawdkmarketingfullrangecreateactivityAPIResponseModel is 创建全场活动 成功返回结果
type AlibabawdkmarketingfullrangecreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_fullrange_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
