package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemSizemappingTemplateUpdate 更新天猫商品尺码表模板
// tmall.item.sizemapping.template.update
//
// 更新天猫商品尺码表模板
func TmallItemSizemappingTemplateUpdate(ctx context.Context, clt *core.SDKClient, req *product.TmallItemSizemappingTemplateUpdateAPIRequest, resp *product.TmallItemSizemappingTemplateUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
