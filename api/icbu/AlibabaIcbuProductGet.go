package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductGet 获得单个商品详情
// alibaba.icbu.product.get
//
// 获取商品详情
func AlibabaIcbuProductGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductGetAPIRequest, resp *icbu.AlibabaIcbuProductGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
