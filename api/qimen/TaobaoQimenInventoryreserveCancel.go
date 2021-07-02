package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenInventoryreserveCancel 库存预占取消接口
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
func TaobaoQimenInventoryreserveCancel(clt *core.SDKClient, req *qimen.TaobaoQimenInventoryreserveCancelAPIRequest, session string) (*qimen.TaobaoQimenInventoryreserveCancelAPIResponse, error) {
	var resp qimen.TaobaoQimenInventoryreserveCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
