package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStockQuery 库存查询接口（多条件）
// taobao.qimen.stock.query
//
// ERP调用奇门的接口,查询商品的库存量
func TaobaoQimenStockQuery(clt *core.SDKClient, req *qimen.TaobaoQimenStockQueryAPIRequest, resp *qimen.TaobaoQimenStockQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
