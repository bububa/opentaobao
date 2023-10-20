package icbuproduct

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// AlibabaIcbuProductInventoryUpdate icbu商品库存更新
// alibaba.icbu.product.inventory.update
//
// 更新库存信息
func AlibabaIcbuProductInventoryUpdate(clt *core.SDKClient, req *icbuproduct.AlibabaIcbuProductInventoryUpdateAPIRequest, resp *icbuproduct.AlibabaIcbuProductInventoryUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
