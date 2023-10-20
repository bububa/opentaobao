package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimeninventoryquery 库存查询接口（多商品）
// taobao.qimen.inventory.query
//
// taobao.qimen.inventory.query
func Taobaoqimeninventoryquery(clt *core.SDKClient, req *qimen.TaobaoqimeninventoryqueryAPIRequest, session string) (*qimen.TaobaoqimeninventoryqueryAPIResponse, error) {
	var resp qimen.TaobaoqimeninventoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
