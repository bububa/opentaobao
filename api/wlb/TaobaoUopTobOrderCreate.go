package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoUopTobOrderCreate ToB仓储发货
// taobao.uop.tob.order.create
//
// ToB仓储发货
func TaobaoUopTobOrderCreate(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoUopTobOrderCreateAPIRequest, resp *wlb.TaobaoUopTobOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
