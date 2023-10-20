package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryWarehouseQuery 分页查询商家仓信息
// taobao.inventory.warehouse.query
//
// 分页查询商家仓信息
func TaobaoInventoryWarehouseQuery(clt *core.SDKClient, req *fenxiao.TaobaoInventoryWarehouseQueryAPIRequest, session string) (*fenxiao.TaobaoInventoryWarehouseQueryAPIResponse, error) {
	var resp fenxiao.TaobaoInventoryWarehouseQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
