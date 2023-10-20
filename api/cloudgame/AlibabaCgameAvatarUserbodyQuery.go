package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCgameAvatarUserbodyQuery 用户Avatar body查询
// alibaba.cgame.avatar.userbody.query
//
// Avatar用户body数据查询
func AlibabaCgameAvatarUserbodyQuery(clt *core.SDKClient, req *cloudgame.AlibabaCgameAvatarUserbodyQueryAPIRequest, resp *cloudgame.AlibabaCgameAvatarUserbodyQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
