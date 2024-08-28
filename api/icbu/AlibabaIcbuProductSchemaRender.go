package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaRender （新）获取商品信息
// alibaba.icbu.product.schema.render
//
// 获取ICBU商品发布的字段填写规则和单个商品对应填写数据，适用于单个商品编辑场景，不包括草稿。
func AlibabaIcbuProductSchemaRender(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaRenderAPIRequest, resp *icbu.AlibabaIcbuProductSchemaRenderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
