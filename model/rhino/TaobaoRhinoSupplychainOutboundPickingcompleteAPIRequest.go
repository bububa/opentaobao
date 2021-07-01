package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest
【WMS005】接收成衣捡配完成通知 API请求
taobao.rhino.supplychain.outbound.pickingcomplete

接收成衣捡配完成通知,WMS005 */
type TaobaoRhinoSupplychainOutboundPickingcompleteAPIRequest struct {
	model.Params
	// 捡配完成消息
	_param0 *PickingCompleteMsg
}

// New
