package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchanneloffline 迎客松影视频道下线
// yunos.tvpubadmin.content.channel.offline
//
// 迎客松影视频道下线
func Yunostvpubadmincontentchanneloffline(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchannelofflineAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchannelofflineAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchannelofflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
