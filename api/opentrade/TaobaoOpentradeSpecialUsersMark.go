package opentrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeSpecialUsersMark 专属下单可购买用户标记
// taobao.opentrade.special.users.mark
//
// 对于专属下单的交易场景，根据openid标记用户可购买商品
func TaobaoOpentradeSpecialUsersMark(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialUsersMarkAPIRequest, resp *opentrade.TaobaoOpentradeSpecialUsersMarkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
