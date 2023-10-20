package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkChannelUserSync 会员同步
// alibaba.wdk.channel.user.sync
//
// 会员同步
func AlibabaWdkChannelUserSync(clt *core.SDKClient, req *wdk.AlibabaWdkChannelUserSyncAPIRequest, resp *wdk.AlibabaWdkChannelUserSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
