package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStockoutConfirm 出库单确认接口
// taobao.qimen.stockout.confirm
//
// 货品出库后，WMS将状态回传给ERP
func TaobaoQimenStockoutConfirm(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenStockoutConfirmAPIRequest, resp *qimen.TaobaoQimenStockoutConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
