package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurBasketMerge 合并购物车
// alibaba.pur.basket.merge
//
// 采购商城接入第三方商家合并购物车接口服务
func AlibabaPurBasketMerge(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurBasketMergeAPIRequest, resp *pur.AlibabaPurBasketMergeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
