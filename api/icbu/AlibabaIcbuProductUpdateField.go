package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductUpdateField 商品按字段更新
// alibaba.icbu.product.update.field
//
// 按字段修改国际站商品，支持询盘商品和在线批发商品，支持英文商品和多语言商品
func AlibabaIcbuProductUpdateField(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductUpdateFieldAPIRequest, resp *icbu.AlibabaIcbuProductUpdateFieldAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
