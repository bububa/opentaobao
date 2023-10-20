package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaAddDraft （新）ICBU商品发布草稿
// alibaba.icbu.product.schema.add.draft
//
// 提供发布ICBU商品草稿的入口
func AlibabaIcbuProductSchemaAddDraft(clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaAddDraftAPIRequest, resp *icbu.AlibabaIcbuProductSchemaAddDraftAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
