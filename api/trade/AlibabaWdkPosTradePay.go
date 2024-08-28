package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkPosTradePay 轻pos品牌营销支付接口
// alibaba.wdk.pos.trade.pay
//
// 轻pos场景，外部商家支付后调用开放平台把支付信息回传给五道口交易
func AlibabaWdkPosTradePay(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaWdkPosTradePayAPIRequest, resp *trade.AlibabaWdkPosTradePayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
