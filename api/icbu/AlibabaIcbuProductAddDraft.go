package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductAddDraft ICBU商品发布草稿接口
// alibaba.icbu.product.add.draft
//
// 发布商品草稿,支持sourcing/一口价商品，支持英文和多种语言原发商品
func AlibabaIcbuProductAddDraft(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductAddDraftAPIRequest, resp *icbu.AlibabaIcbuProductAddDraftAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
