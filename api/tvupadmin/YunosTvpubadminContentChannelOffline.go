package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChannelOffline 迎客松影视频道下线
// yunos.tvpubadmin.content.channel.offline
//
// 迎客松影视频道下线
func YunosTvpubadminContentChannelOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChannelOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentChannelOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
