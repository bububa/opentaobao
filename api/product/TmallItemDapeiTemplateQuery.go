package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemDapeiTemplateQuery 搭配查询接口
// tmall.item.dapei.template.query
//
// 根据条件获取搭配内容
func TmallItemDapeiTemplateQuery(ctx context.Context, clt *core.SDKClient, req *product.TmallItemDapeiTemplateQueryAPIRequest, resp *product.TmallItemDapeiTemplateQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
