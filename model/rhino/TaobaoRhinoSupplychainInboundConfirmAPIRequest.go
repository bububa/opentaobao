package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRhinoSupplychainInboundConfirmAPIRequest
WMS003成衣入库确认 API请求
taobao.rhino.supplychain.inbound.confirm

【WMS003】【同步成衣入库完成信息】 */
type TaobaoRhinoSupplychainInboundConfirmAPIRequest struct {
	model.Params
	// 入库单确认对象
	_clothingInboundConfirm *ClothingInboundConfirmDto
}

// New
