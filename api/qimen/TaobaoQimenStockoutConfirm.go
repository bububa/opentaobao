package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStockoutConfirm 出库单确认接口
// taobao.qimen.stockout.confirm
//
// 货品出库后，WMS将状态回传给ERP
func TaobaoQimenStockoutConfirm(clt *core.SDKClient, req *qimen.TaobaoQimenStockoutConfirmAPIRequest, session string) (*qimen.TaobaoQimenStockoutConfirmAPIResponse, error) {
	var resp qimen.TaobaoQimenStockoutConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
