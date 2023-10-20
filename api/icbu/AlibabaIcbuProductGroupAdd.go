package icbu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductGroupAdd 增加商品分组
// alibaba.icbu.product.group.add
//
// 增加商品分组
func AlibabaIcbuProductGroupAdd(clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupAddAPIRequest, resp *icbu.AlibabaIcbuProductGroupAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
