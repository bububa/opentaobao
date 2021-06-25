package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
会员同步 APIResponse
alibaba.wdk.channel.user.sync

会员同步
*/
type AlibabaWdkChannelUserSyncAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkChannelUserSyncResponse `json:"alibaba_wdk_channel_user_sync_response,omitempty"`
}

type AlibabaWdkChannelUserSyncResponse struct {

    // 返回内容
    ApiResult   *AlibabaWdkChannelUserSyncApiResult `json:"api_result,omitempty"`

}
