package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsReturnBillGet 获取销退收货信息
// taobao.wlb.wms.return.bill.get
//
// 通过订单号获取单个销退单收货信息。
func TaobaoWlbWmsReturnBillGet(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsReturnBillGetAPIRequest, resp *wms.TaobaoWlbWmsReturnBillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
