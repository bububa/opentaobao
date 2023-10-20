package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse 活动数据同步 API返回值
// alibaba.wdk.marketing.open.darunfa.activity.sync
//
// 大润发活动数据同步
type AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponseModel).Reset()
}

// AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponseModel is 活动数据同步 成功返回结果
type AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_darunfa_activity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse)
	},
}

// GetAlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse
func GetAlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse() *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse {
	return poolAlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse.Get().(*AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse)
}

// ReleaseAlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse 将 AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse(v *AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse.Put(v)
}
