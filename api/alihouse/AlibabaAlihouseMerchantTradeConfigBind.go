package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMerchantTradeConfigBind 交易场景绑定
// alibaba.alihouse.merchant.trade.config.bind
//
// 交易场景绑定
func AlibabaAlihouseMerchantTradeConfigBind(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseMerchantTradeConfigBindAPIRequest, resp *alihouse.AlibabaAlihouseMerchantTradeConfigBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
