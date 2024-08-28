package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductAddSchemaGet 产品发布规则获取接口
// tmall.product.add.schema.get
//
// 获取用户发布产品的规则
func TmallProductAddSchemaGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallProductAddSchemaGetAPIRequest, resp *tbitem.TmallProductAddSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
