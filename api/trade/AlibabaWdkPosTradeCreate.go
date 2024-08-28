package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkPosTradeCreate 轻pos品牌营销下单接口
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
func AlibabaWdkPosTradeCreate(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaWdkPosTradeCreateAPIRequest, resp *trade.AlibabaWdkPosTradeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
