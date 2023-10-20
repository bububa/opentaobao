package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// AlibabaCloudgameOpenidQuery 咖哒用户ID查询
// alibaba.cloudgame.openid.query
//
// 云游戏业务需要提供查询用户信息的TOP能力
func AlibabaCloudgameOpenidQuery(clt *core.SDKClient, req *cloudgame.AlibabaCloudgameOpenidQueryAPIRequest, resp *cloudgame.AlibabaCloudgameOpenidQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
