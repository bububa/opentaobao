package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoOpentradeCustomizationRefundEnable 定制订单设置允许仅退款
// taobao.opentrade.customization.refund.enable
//
// 定制订单设置允许仅退款
func TaobaoOpentradeCustomizationRefundEnable(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoOpentradeCustomizationRefundEnableAPIRequest, resp *trade.TaobaoOpentradeCustomizationRefundEnableAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
