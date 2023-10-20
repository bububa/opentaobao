package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentChannelQuery 迎客松影视频道查询
// yunos.tvpubadmin.content.channel.query
//
// 迎客松影视频道查询
func YunosTvpubadminContentChannelQuery(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentChannelQueryAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentChannelQueryAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentChannelQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
