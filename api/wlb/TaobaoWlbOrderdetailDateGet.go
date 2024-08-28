package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderdetailDateGet 按照日期范围查询物流订单详情
// taobao.wlb.orderdetail.date.get
//
// 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情
func TaobaoWlbOrderdetailDateGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbOrderdetailDateGetAPIRequest, resp *wlb.TaobaoWlbOrderdetailDateGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
