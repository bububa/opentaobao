package crm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmShopvipCancel 卖家取消店铺vip的优惠
// taobao.crm.shopvip.cancel
//
// 此接口用于取消VIP优惠
func TaobaoCrmShopvipCancel(ctx context.Context, clt *core.SDKClient, req *crm.TaobaoCrmShopvipCancelAPIRequest, resp *crm.TaobaoCrmShopvipCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
