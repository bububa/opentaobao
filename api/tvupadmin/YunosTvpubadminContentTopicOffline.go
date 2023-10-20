package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentTopicOffline 迎客松专题下线
// yunos.tvpubadmin.content.topic.offline
//
// 迎客松专题下线
func YunosTvpubadminContentTopicOffline(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentTopicOfflineAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentTopicOfflineAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentTopicOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
