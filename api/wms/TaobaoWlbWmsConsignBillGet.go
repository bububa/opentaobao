package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsConsignBillGet 获取销售订单发货信息
// taobao.wlb.wms.consign.bill.get
//
// 获取销售订单发货信息
func TaobaoWlbWmsConsignBillGet(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsConsignBillGetAPIRequest, resp *wms.TaobaoWlbWmsConsignBillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
