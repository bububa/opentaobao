package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmincontentchannelquery 迎客松影视频道查询
// yunos.tvpubadmin.content.channel.query
//
// 迎客松影视频道查询
func Yunostvpubadmincontentchannelquery(clt *core.SDKClient, req *tvupadmin.YunostvpubadmincontentchannelqueryAPIRequest, session string) (*tvupadmin.YunostvpubadmincontentchannelqueryAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmincontentchannelqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
