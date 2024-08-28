package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelFutureSoftmodify 未来酒店信息下发
// taobao.xhotel.future.softmodify
//
// 未来酒店信息下发，包含PMS订单查询和自助入住
func TaobaoXhotelFutureSoftmodify(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelFutureSoftmodifyAPIRequest, resp *xhotelonlineorder.TaobaoXhotelFutureSoftmodifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
