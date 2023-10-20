package lstvending

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstvending"
)

// AlibabaLstVendingShippingCallback 售货机出货回传接口
// alibaba.lst.vending.shipping.callback
//
// 零售通自动售货机商品出货回传接口，同步商品出库最新状态。
func AlibabaLstVendingShippingCallback(clt *core.SDKClient, req *lstvending.AlibabaLstVendingShippingCallbackAPIRequest, resp *lstvending.AlibabaLstVendingShippingCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
