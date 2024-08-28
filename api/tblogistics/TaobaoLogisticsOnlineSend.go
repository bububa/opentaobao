package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsOnlineSend 在线订单发货处理（支持货到付款）
// taobao.logistics.online.send
//
// 用户调用该接口可实现在线订单发货（支持货到付款）&lt;br/&gt;调用该接口实现在线下单发货，有两种情况：&lt;br&gt;&lt;br/&gt;&lt;font color=&#39;red&#39;&gt;如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。&lt;br&gt;&lt;br/&gt;如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。&lt;/font&gt;
func TaobaoLogisticsOnlineSend(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsOnlineSendAPIRequest, resp *tblogistics.TaobaoLogisticsOnlineSendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
