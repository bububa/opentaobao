package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstockquery 库存查询接口（多条件）
// taobao.qimen.stock.query
//
// ERP调用奇门的接口,查询商品的库存量
func Taobaoqimenstockquery(clt *core.SDKClient, req *qimen.TaobaoqimenstockqueryAPIRequest, session string) (*qimen.TaobaoqimenstockqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenstockqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
