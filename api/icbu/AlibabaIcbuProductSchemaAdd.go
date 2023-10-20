package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaAdd （新）商品发布新接口
// alibaba.icbu.product.schema.add
//
// 提供发布ICBU商品的入口
func AlibabaIcbuProductSchemaAdd(clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaAddAPIRequest, resp *icbu.AlibabaIcbuProductSchemaAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
