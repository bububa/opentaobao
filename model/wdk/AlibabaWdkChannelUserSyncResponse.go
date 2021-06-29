package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员同步 API返回值 
alibaba.wdk.channel.user.sync

会员同步
*/
type AlibabaWdkChannelUserSyncAPIResponse struct {
    model.CommonResponse
    AlibabaWdkChannelUserSyncResponse
}

// 会员同步 成功返回结果
type AlibabaWdkChannelUserSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_channel_user_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回内容
    ApiResult   *AlibabaWdkChannelUserSyncApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
