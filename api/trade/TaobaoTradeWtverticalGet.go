package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradeWtverticalGet 网厅垂直信息查询接口
// taobao.trade.wtvertical.get
//
// 网厅订单垂直信息的查询
func TaobaoTradeWtverticalGet(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoTradeWtverticalGetAPIRequest, resp *trade.TaobaoTradeWtverticalGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
