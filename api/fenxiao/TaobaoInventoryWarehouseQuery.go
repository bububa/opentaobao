package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryWarehouseQuery 分页查询商家仓信息
// taobao.inventory.warehouse.query
//
// 分页查询商家仓信息
func TaobaoInventoryWarehouseQuery(clt *core.SDKClient, req *fenxiao.TaobaoInventoryWarehouseQueryAPIRequest, resp *fenxiao.TaobaoInventoryWarehouseQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
