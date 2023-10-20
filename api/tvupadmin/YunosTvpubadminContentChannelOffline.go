package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChannelOffline 迎客松影视频道下线
// yunos.tvpubadmin.content.channel.offline
//
// 迎客松影视频道下线
func YunosTvpubadminContentChannelOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChannelOfflineAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentChannelOfflineAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentChannelOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
