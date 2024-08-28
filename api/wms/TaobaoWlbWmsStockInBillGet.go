package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsStockInBillGet 获取入库单信息
// taobao.wlb.wms.stock.in.bill.get
//
// 获取入库单信息
func TaobaoWlbWmsStockInBillGet(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsStockInBillGetAPIRequest, resp *wms.TaobaoWlbWmsStockInBillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
