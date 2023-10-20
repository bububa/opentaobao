package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimeninventoryreservecancel 库存预占取消接口
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
func Taobaoqimeninventoryreservecancel(clt *core.SDKClient, req *qimen.TaobaoqimeninventoryreservecancelAPIRequest, session string) (*qimen.TaobaoqimeninventoryreservecancelAPIResponse, error) {
	var resp qimen.TaobaoqimeninventoryreservecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
