package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoOrderQrcode 根据商品id列表获取可扫描下单二维码
// tmall.aliauto.order.qrcode
//
// 根据商品id列表获取可扫描下单二维码
func TmallAliautoOrderQrcode(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallAliautoOrderQrcodeAPIRequest, resp *tmallcar.TmallAliautoOrderQrcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
