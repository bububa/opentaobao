package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkPosTradeReverse 轻pos品牌营销退款接口
// alibaba.wdk.pos.trade.reverse
//
// 轻pos品牌营销场景，商家调用退款接口
func AlibabaWdkPosTradeReverse(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaWdkPosTradeReverseAPIRequest, resp *trade.AlibabaWdkPosTradeReverseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
