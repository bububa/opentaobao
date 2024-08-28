package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaRenderDraft （新）渲染草稿商品数据
// alibaba.icbu.product.schema.render.draft
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个草稿商品编辑场景，
func AlibabaIcbuProductSchemaRenderDraft(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaRenderDraftAPIRequest, resp *icbu.AlibabaIcbuProductSchemaRenderDraftAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
