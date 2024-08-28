package tbrefund

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// TaobaoRefundStatusGet 查询退款状态
// taobao.refund.status.get
//
// 根据主订单或者子订单id查询退款状态,入参传入主订单或者子订单号1.如果当传入子订单时，返回子订单最后一笔退款单的状态,如果子订单申请退款退款返回空list.2.如果传传入主订单，则返回主订单下所有子订单的最后一笔退款单状态，如果对应的子订单没有生成退款单，则对应子订单对应数据返回。
func TaobaoRefundStatusGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRefundStatusGetAPIRequest, resp *tbrefund.TaobaoRefundStatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
