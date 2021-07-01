package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenInventoryreserveCancelAPIRequest
库存预占取消接口 API请求
taobao.qimen.inventoryreserve.cancel

库存预占取消 */
type TaobaoQimenInventoryreserveCancelAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenInventoryreserveCancelRequest
}

// New
