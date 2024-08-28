package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TmallDeviceItemPromotionGet 智能硬件上商品优惠获取
// tmall.device.item.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
func TmallDeviceItemPromotionGet(ctx context.Context, clt *core.SDKClient, req *iot.TmallDeviceItemPromotionGetAPIRequest, resp *iot.TmallDeviceItemPromotionGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
