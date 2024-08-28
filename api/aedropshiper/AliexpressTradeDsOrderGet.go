package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressTradeDsOrderGet 买家查询订单详情
// aliexpress.trade.ds.order.get
//
// 买家查询订单详情，用于dropshipper
func AliexpressTradeDsOrderGet(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressTradeDsOrderGetAPIRequest, resp *aedropshiper.AliexpressTradeDsOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
