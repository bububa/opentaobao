package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransRefundCreate 阿信创建退款单
// taobao.alitrip.axin.trans.refund.create
//
// 阿信供销平台-创建退款单服务
func TaobaoAlitripAxinTransRefundCreate(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransRefundCreateAPIRequest, resp *axintrade.TaobaoAlitripAxinTransRefundCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
