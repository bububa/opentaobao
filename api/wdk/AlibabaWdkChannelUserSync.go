package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
会员同步 
alibaba.wdk.channel.user.sync

会员同步
*/
func AlibabaWdkChannelUserSync(clt *core.SDKClient, req *wdk.AlibabaWdkChannelUserSyncRequest, session string) (*wdk.AlibabaWdkChannelUserSyncResponse, error) {
    var resp wdk.AlibabaWdkChannelUserSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
