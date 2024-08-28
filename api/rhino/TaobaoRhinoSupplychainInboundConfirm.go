package rhino

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// TaobaoRhinoSupplychainInboundConfirm WMS003成衣入库确认
// taobao.rhino.supplychain.inbound.confirm
//
// 【WMS003】【同步成衣入库完成信息】
func TaobaoRhinoSupplychainInboundConfirm(ctx context.Context, clt *core.SDKClient, req *rhino.TaobaoRhinoSupplychainInboundConfirmAPIRequest, resp *rhino.TaobaoRhinoSupplychainInboundConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
