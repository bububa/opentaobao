package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallProductSpecAdd 添加产品规格
// tmall.product.spec.add
//
// 增加产品规格
func TmallProductSpecAdd(ctx context.Context, clt *core.SDKClient, req *product.TmallProductSpecAddAPIRequest, resp *product.TmallProductSpecAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
