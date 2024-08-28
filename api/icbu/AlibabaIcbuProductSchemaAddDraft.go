package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaAddDraft （新）ICBU商品发布草稿
// alibaba.icbu.product.schema.add.draft
//
// 提供发布ICBU商品草稿的入口
func AlibabaIcbuProductSchemaAddDraft(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaAddDraftAPIRequest, resp *icbu.AlibabaIcbuProductSchemaAddDraftAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
