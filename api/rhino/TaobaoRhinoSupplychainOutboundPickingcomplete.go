package rhino

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

/* TaobaoRhinoSupplychainOutboundPickingcomplete
【WMS005】接收成衣捡配完成通知
taobao.rhino.supplychain.outbound.pickingcomplete

接收成衣捡配完成通知,WMS005 */
func TaobaoRhinoSupplychainOutboundPickingcomplete(clt *core.SDKClient, req *rhino.TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest, session string) (*rhino.TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse, error) {
	var resp rhino.TaobaoRhinoSupplychainOutboundPickingcompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
