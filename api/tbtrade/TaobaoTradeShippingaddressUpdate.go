package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradeShippingaddressUpdate 更改交易的收货地址
// taobao.trade.shippingaddress.update
//
// 只能更新一笔交易里面的买家收货地址 &lt;br/&gt;只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 &lt;br/&gt;更新后的发货地址可以通过taobao.trade.fullinfo.get查到 &lt;br/&gt;参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
func TaobaoTradeShippingaddressUpdate(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTradeShippingaddressUpdateAPIRequest, resp *tbtrade.TaobaoTradeShippingaddressUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
