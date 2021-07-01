package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStockQueryAPIRequest
库存查询接口（多条件） API请求
taobao.qimen.stock.query

ERP调用奇门的接口,查询商品的库存量 */
type TaobaoQimenStockQueryAPIRequest struct {
	model.Params
	//
	_request *StockQueryRequest
}

// New
