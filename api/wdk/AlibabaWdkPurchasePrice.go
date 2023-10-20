package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkPurchasePrice rt回传采购价
// alibaba.wdk.purchase.price
//
// 猫超共享库存项目-rt回传采购价
func AlibabaWdkPurchasePrice(clt *core.SDKClient, req *wdk.AlibabaWdkPurchasePriceAPIRequest, resp *wdk.AlibabaWdkPurchasePriceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
