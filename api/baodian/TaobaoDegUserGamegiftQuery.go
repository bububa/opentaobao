package baodian

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// TaobaoDegUserGamegiftQuery 用户数娱游戏礼包查询
// taobao.deg.user.gamegift.query
//
// 查询用户数娱礼包列表
func TaobaoDegUserGamegiftQuery(ctx context.Context, clt *core.SDKClient, req *baodian.TaobaoDegUserGamegiftQueryAPIRequest, resp *baodian.TaobaoDegUserGamegiftQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
