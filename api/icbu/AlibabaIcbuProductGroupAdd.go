package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductGroupAdd 增加商品分组
// alibaba.icbu.product.group.add
//
// 增加商品分组
func AlibabaIcbuProductGroupAdd(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupAddAPIRequest, resp *icbu.AlibabaIcbuProductGroupAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
