package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderConsign 物流宝订单已发货通知接口
// taobao.wlb.order.consign
//
// 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货
func TaobaoWlbOrderConsign(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbOrderConsignAPIRequest, resp *wlb.TaobaoWlbOrderConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
