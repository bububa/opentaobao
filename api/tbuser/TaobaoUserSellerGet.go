package tbuser

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserSellerGet 查询卖家用户信息
// taobao.user.seller.get
//
// 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
func TaobaoUserSellerGet(ctx context.Context, clt *core.SDKClient, req *tbuser.TaobaoUserSellerGetAPIRequest, resp *tbuser.TaobaoUserSellerGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
