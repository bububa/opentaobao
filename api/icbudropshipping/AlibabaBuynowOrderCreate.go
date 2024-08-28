package icbudropshipping

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// AlibabaBuynowOrderCreate 阿里巴巴买家buynow下单接口
// alibaba.buynow.order.create
//
// 阿里巴巴买家下单接口
func AlibabaBuynowOrderCreate(ctx context.Context, clt *core.SDKClient, req *icbudropshipping.AlibabaBuynowOrderCreateAPIRequest, resp *icbudropshipping.AlibabaBuynowOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
