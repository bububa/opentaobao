package baodian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// TaobaoDegUserGamegiftQuery 用户数娱游戏礼包查询
// taobao.deg.user.gamegift.query
//
// 查询用户数娱礼包列表
func TaobaoDegUserGamegiftQuery(clt *core.SDKClient, req *baodian.TaobaoDegUserGamegiftQueryAPIRequest, resp *baodian.TaobaoDegUserGamegiftQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
