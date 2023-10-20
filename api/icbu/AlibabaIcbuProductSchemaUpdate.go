package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductSchemaUpdate （新）商品发布增量更新接口
// alibaba.icbu.product.schema.update
//
// 商品更新接口，方式为增量更新，只更新传入的字段
func AlibabaIcbuProductSchemaUpdate(clt *core.SDKClient, req *icbu.AlibabaIcbuProductSchemaUpdateAPIRequest, resp *icbu.AlibabaIcbuProductSchemaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
