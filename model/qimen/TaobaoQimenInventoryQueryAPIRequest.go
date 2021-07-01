package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenInventoryQueryAPIRequest
库存查询接口（多商品） API请求
taobao.qimen.inventory.query

ERP调用奇门的接口,查询商品的库存量 */
type TaobaoQimenInventoryQueryAPIRequest struct {
	model.Params
	//
	_request *InventoryQueryRequest
}

// New
