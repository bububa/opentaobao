package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallProductMatchSchemaGet 获取匹配产品规则
// tmall.product.match.schema.get
//
// ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
func TmallProductMatchSchemaGet(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallProductMatchSchemaGetAPIRequest, resp *tbitem.TmallProductMatchSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
