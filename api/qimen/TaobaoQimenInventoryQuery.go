package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventoryQuery 库存查询接口（多商品）
// taobao.qimen.inventory.query
//
// taobao.qimen.inventory.query
func TaobaoQimenInventoryQuery(clt *core.SDKClient, req *qimen.TaobaoQimenInventoryQueryAPIRequest, session string) (*qimen.TaobaoQimenInventoryQueryAPIResponse, error) {
	var resp qimen.TaobaoQimenInventoryQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
