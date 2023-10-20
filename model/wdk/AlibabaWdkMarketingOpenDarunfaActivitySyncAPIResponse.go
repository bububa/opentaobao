package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingopendarunfaactivitysyncAPIResponse 活动数据同步 API返回值
// alibaba.wdk.marketing.open.darunfa.activity.sync
//
// 大润发活动数据同步
type AlibabawdkmarketingopendarunfaactivitysyncAPIResponse struct {
	model.CommonResponse
	AlibabawdkmarketingopendarunfaactivitysyncAPIResponseModel
}

// AlibabawdkmarketingopendarunfaactivitysyncAPIResponseModel is 活动数据同步 成功返回结果
type AlibabawdkmarketingopendarunfaactivitysyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_darunfa_activity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
