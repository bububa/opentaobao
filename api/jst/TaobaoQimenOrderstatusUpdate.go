package jst

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoQimenOrderstatusUpdate 订单状态更新接口
// taobao.qimen.orderstatus.update
//
// 星盘和ISV，可以通过此接口，来更新订单状态。此接口应用于使用阿里星盘分单，且使用商家系统（非阿里掌柜）接单/拒单的模式下更新订单状态。
func TaobaoQimenOrderstatusUpdate(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoQimenOrderstatusUpdateAPIRequest, resp *jst.TaobaoQimenOrderstatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
