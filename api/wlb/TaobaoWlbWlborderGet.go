package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbWlborderGet 根据物流宝订单编号查询物流宝订单概要信息
// taobao.wlb.wlborder.get
//
// 根据物流宝订单编号查询物流宝订单概要信息
func TaobaoWlbWlborderGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbWlborderGetAPIRequest, resp *wlb.TaobaoWlbWlborderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
