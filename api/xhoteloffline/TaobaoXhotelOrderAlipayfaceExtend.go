package xhoteloffline

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhoteloffline"
)

// TaobaoXhotelOrderAlipayfaceExtend 信用住订单延住接口
// taobao.xhotel.order.alipayface.extend
//
// 信用住订单延住接口，用于将已有的信用住支付订单checkOut和订单金额等更新
func TaobaoXhotelOrderAlipayfaceExtend(ctx context.Context, clt *core.SDKClient, req *xhoteloffline.TaobaoXhotelOrderAlipayfaceExtendAPIRequest, resp *xhoteloffline.TaobaoXhotelOrderAlipayfaceExtendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
