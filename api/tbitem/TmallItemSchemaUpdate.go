package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemSchemaUpdate 天猫根据规则编辑商品
// tmall.item.schema.update
//
// 天猫根据规则编辑商品
func TmallItemSchemaUpdate(ctx context.Context, clt *core.SDKClient, req *tbitem.TmallItemSchemaUpdateAPIRequest, resp *tbitem.TmallItemSchemaUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
