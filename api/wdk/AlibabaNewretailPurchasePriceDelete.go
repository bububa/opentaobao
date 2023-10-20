package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaNewretailPurchasePriceDelete 共享库存 商户删除采购价
// alibaba.newretail.purchase.price.delete
//
// 共享库存 商户删除采购价
func AlibabaNewretailPurchasePriceDelete(clt *core.SDKClient, req *wdk.AlibabaNewretailPurchasePriceDeleteAPIRequest, resp *wdk.AlibabaNewretailPurchasePriceDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
