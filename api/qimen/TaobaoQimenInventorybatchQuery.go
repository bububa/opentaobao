package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimeninventorybatchquery 商品单仓批次库存查询接口
// taobao.qimen.inventorybatch.query
//
// ERP 通过该接口查询指定商品的单仓批次库存
func Taobaoqimeninventorybatchquery(clt *core.SDKClient, req *qimen.TaobaoqimeninventorybatchqueryAPIRequest, session string) (*qimen.TaobaoqimeninventorybatchqueryAPIResponse, error) {
	var resp qimen.TaobaoqimeninventorybatchqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
