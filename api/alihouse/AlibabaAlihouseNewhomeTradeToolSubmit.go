package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeToolSubmit 交易工具信息上翻
// alibaba.alihouse.newhome.trade.tool.submit
//
// 交易工具信息上翻
func AlibabaAlihouseNewhomeTradeToolSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeToolSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeTradeToolSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
