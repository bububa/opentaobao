package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSchemaAdd 天猫根据规则发布商品
// tmall.item.schema.add
//
// 天猫TopSchema发布商品。
func TmallItemSchemaAdd(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemSchemaAddAPIRequest, resp *tbitem.TmallItemSchemaAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
