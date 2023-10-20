package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductAdd 发布产品
// alibaba.icbu.product.add
//
// 发布商品,支持sourcing/一口价商品，支持英文和多种语言原发商品
func AlibabaIcbuProductAdd(clt *core.SDKClient, req *icbu.AlibabaIcbuProductAddAPIRequest, resp *icbu.AlibabaIcbuProductAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
