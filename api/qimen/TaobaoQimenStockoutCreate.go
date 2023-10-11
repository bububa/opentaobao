package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStockoutCreate 出库单创建接口
// taobao.qimen.stockout.create
//
// taobao.qimen.returnorder.create
func TaobaoQimenStockoutCreate(clt *core.SDKClient, req *qimen.TaobaoQimenStockoutCreateAPIRequest, session string) (*qimen.TaobaoQimenStockoutCreateAPIResponse, error) {
	var resp qimen.TaobaoQimenStockoutCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
