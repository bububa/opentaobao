package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderAlipayfaceCancelsettle 信用住订单取消结算接口
// taobao.xhotel.order.alipayface.cancelsettle
//
// 信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。
func TaobaoXhotelOrderAlipayfaceCancelsettle(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderAlipayfaceCancelsettleAPIRequest, resp *xhotelonlineorder.TaobaoXhotelOrderAlipayfaceCancelsettleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
