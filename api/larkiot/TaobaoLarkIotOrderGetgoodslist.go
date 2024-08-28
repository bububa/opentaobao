package larkiot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// TaobaoLarkIotOrderGetgoodslist iot渠道获取卖品信息
// taobao.lark.iot.order.getgoodslist
//
// iot无人超市服务商通过接口获取影院的可售卖品数据
func TaobaoLarkIotOrderGetgoodslist(ctx context.Context, clt *core.SDKClient, req *larkiot.TaobaoLarkIotOrderGetgoodslistAPIRequest, resp *larkiot.TaobaoLarkIotOrderGetgoodslistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
