package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryWarehouseManage 创建商家仓或者更新商家仓信息
// taobao.inventory.warehouse.manage
//
// 创建商家仓或者更新商家仓信息
func TaobaoInventoryWarehouseManage(clt *core.SDKClient, req *fenxiao.TaobaoInventoryWarehouseManageAPIRequest, session string) (*fenxiao.TaobaoInventoryWarehouseManageAPIResponse, error) {
	var resp fenxiao.TaobaoInventoryWarehouseManageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
