package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaGet （新）ICBU商品发布schema接口
// alibaba.icbu.product.schema.get
//
// 获取ICBU商品发布的页面规则和填写字段，适用于新发商品
func AlibabaIcbuProductSchemaGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaGetAPIRequest, resp *icbu.AlibabaIcbuProductSchemaGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
