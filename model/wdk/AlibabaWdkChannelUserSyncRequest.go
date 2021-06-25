package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员同步 APIRequest
alibaba.wdk.channel.user.sync

会员同步
*/
type AlibabaWdkChannelUserSyncRequest struct {
    model.Params

    // 会员信息
    userSyncInfo   *UserSyncInfo 

}

func NewAlibabaWdkChannelUserSyncRequest() *AlibabaWdkChannelUserSyncRequest{
    return &AlibabaWdkChannelUserSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkChannelUserSyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.user.sync"
}

func (r AlibabaWdkChannelUserSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkChannelUserSyncRequest) SetUserSyncInfo(userSyncInfo *UserSyncInfo) error {
    r.userSyncInfo = userSyncInfo
    r.Set("user_sync_info", userSyncInfo)
    return nil
}

func (r AlibabaWdkChannelUserSyncRequest) GetUserSyncInfo() *UserSyncInfo {
    return r.userSyncInfo
}

