package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuCategorySchemaLevelGet (新)层级属性获取
// alibaba.icbu.category.schema.level.get
//
// 将表单中层级属性的子属性返回
func AlibabaIcbuCategorySchemaLevelGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuCategorySchemaLevelGetAPIRequest, resp *icbu.AlibabaIcbuCategorySchemaLevelGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
