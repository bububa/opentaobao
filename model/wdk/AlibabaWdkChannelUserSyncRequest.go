package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员同步 API请求
alibaba.wdk.channel.user.sync

会员同步
*/
type AlibabaWdkChannelUserSyncRequest struct {
    model.Params
    // 会员信息
    userSyncInfo   *UserSyncInfo
}

// 初始化AlibabaWdkChannelUserSyncRequest对象
func NewAlibabaWdkChannelUserSyncRequest() *AlibabaWdkChannelUserSyncRequest{
    return &AlibabaWdkChannelUserSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelUserSyncRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.user.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelUserSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserSyncInfo Setter
// 会员信息
func (r *AlibabaWdkChannelUserSyncRequest) SetUserSyncInfo(userSyncInfo *UserSyncInfo) error {
    r.userSyncInfo = userSyncInfo
    r.Set("user_sync_info", userSyncInfo)
    return nil
}

// UserSyncInfo Getter
func (r AlibabaWdkChannelUserSyncRequest) GetUserSyncInfo() *UserSyncInfo {
    return r.userSyncInfo
}
