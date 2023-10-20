package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerVendorTradePurchase 查看购买人的订单记录以及授权时间
// alibaba.seller.vendor.trade.purchase
//
// 查看购买人的订单记录以及授权时间
func AlibabaSellerVendorTradePurchase(clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorTradePurchaseAPIRequest, resp *icbuseller.AlibabaSellerVendorTradePurchaseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
