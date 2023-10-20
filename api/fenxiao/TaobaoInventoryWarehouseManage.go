package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoinventorywarehousemanage 创建商家仓或者更新商家仓信息
// taobao.inventory.warehouse.manage
//
// 创建商家仓或者更新商家仓信息
func Taobaoinventorywarehousemanage(clt *core.SDKClient, req *fenxiao.TaobaoinventorywarehousemanageAPIRequest, session string) (*fenxiao.TaobaoinventorywarehousemanageAPIResponse, error) {
	var resp fenxiao.TaobaoinventorywarehousemanageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
