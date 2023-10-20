package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTopicOffline 迎客松专题下线
// yunos.tvpubadmin.content.topic.offline
//
// 迎客松专题下线
func YunosTvpubadminContentTopicOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTopicOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminContentTopicOfflineAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
