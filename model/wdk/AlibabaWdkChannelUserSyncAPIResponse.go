package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelUserSyncAPIResponse 会员同步 API返回值
// alibaba.wdk.channel.user.sync
//
// 会员同步
type AlibabaWdkChannelUserSyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelUserSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkChannelUserSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkChannelUserSyncAPIResponseModel).Reset()
}

// AlibabaWdkChannelUserSyncAPIResponseModel is 会员同步 成功返回结果
type AlibabaWdkChannelUserSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_user_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回内容
	ApiResult *AlibabaWdkChannelUserSyncApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkChannelUserSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkChannelUserSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelUserSyncAPIResponse)
	},
}

// GetAlibabaWdkChannelUserSyncAPIResponse 从 sync.Pool 获取 AlibabaWdkChannelUserSyncAPIResponse
func GetAlibabaWdkChannelUserSyncAPIResponse() *AlibabaWdkChannelUserSyncAPIResponse {
	return poolAlibabaWdkChannelUserSyncAPIResponse.Get().(*AlibabaWdkChannelUserSyncAPIResponse)
}

// ReleaseAlibabaWdkChannelUserSyncAPIResponse 将 AlibabaWdkChannelUserSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkChannelUserSyncAPIResponse(v *AlibabaWdkChannelUserSyncAPIResponse) {
	v.Reset()
	poolAlibabaWdkChannelUserSyncAPIResponse.Put(v)
}
