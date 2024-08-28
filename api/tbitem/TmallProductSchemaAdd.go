package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductSchemaAdd 使用Schema文件发布一个产品
// tmall.product.schema.add
//
// Schema体系发布一个产品
func TmallProductSchemaAdd(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallProductSchemaAddAPIRequest, resp *tbitem.TmallProductSchemaAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
