package opentrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeSpecialUsersQuery 专属下单标记信息查询
// taobao.opentrade.special.users.query
//
// 专属下单标记信息查询
func TaobaoOpentradeSpecialUsersQuery(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialUsersQueryAPIRequest, resp *opentrade.TaobaoOpentradeSpecialUsersQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
