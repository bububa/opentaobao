package rhino

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// TaobaoRhinoSupplychainOutboundPickingcomplete 【WMS005】接收成衣捡配完成通知
// taobao.rhino.supplychain.outbound.pickingcomplete
//
// 接收成衣捡配完成通知,WMS005
func TaobaoRhinoSupplychainOutboundPickingcomplete(ctx context.Context, clt *core.SDKClient, req *rhino.TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest, resp *rhino.TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
