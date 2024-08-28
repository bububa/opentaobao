package icbudropshipping

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaDropshippingTokenCreate 国际站dropshipping 选品token 创建
// alibaba.dropshipping.token.create
//
// 国际站dropshipping 选品token 创建，用于让买家有权限访问我们指定的 商品场馆
func AlibabaDropshippingTokenCreate(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaDropshippingTokenCreateAPIRequest, resp *icbudropshipping.AlibabaDropshippingTokenCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
