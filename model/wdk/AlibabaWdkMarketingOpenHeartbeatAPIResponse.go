package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenHeartbeatAPIResponse 心跳服务【10s一次】 API返回值
// alibaba.wdk.marketing.open.heartbeat
//
// 商家数据同步心跳服务
type AlibabaWdkMarketingOpenHeartbeatAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingOpenHeartbeatAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenHeartbeatAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingOpenHeartbeatAPIResponseModel).Reset()
}

// AlibabaWdkMarketingOpenHeartbeatAPIResponseModel is 心跳服务【10s一次】 成功返回结果
type AlibabaWdkMarketingOpenHeartbeatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_heartbeat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenHeartbeatAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingOpenHeartbeatAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingOpenHeartbeatAPIResponse)
	},
}

// GetAlibabaWdkMarketingOpenHeartbeatAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingOpenHeartbeatAPIResponse
func GetAlibabaWdkMarketingOpenHeartbeatAPIResponse() *AlibabaWdkMarketingOpenHeartbeatAPIResponse {
	return poolAlibabaWdkMarketingOpenHeartbeatAPIResponse.Get().(*AlibabaWdkMarketingOpenHeartbeatAPIResponse)
}

// ReleaseAlibabaWdkMarketingOpenHeartbeatAPIResponse 将 AlibabaWdkMarketingOpenHeartbeatAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingOpenHeartbeatAPIResponse(v *AlibabaWdkMarketingOpenHeartbeatAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingOpenHeartbeatAPIResponse.Put(v)
}
