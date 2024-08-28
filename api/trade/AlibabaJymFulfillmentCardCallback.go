package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaJymFulfillmentCardCallback 外部商家卡密结果回调
// alibaba.jym.fulfillment.card.callback
//
// 外部商家卡密结果回调
func AlibabaJymFulfillmentCardCallback(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaJymFulfillmentCardCallbackAPIRequest, resp *trade.AlibabaJymFulfillmentCardCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
