package icbuseller

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerVendorTradePurchase 查看购买人的订单记录以及授权时间
// alibaba.seller.vendor.trade.purchase
//
// 查看购买人的订单记录以及授权时间
func AlibabaSellerVendorTradePurchase(ctx context.Context, clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorTradePurchaseAPIRequest, resp *icbuseller.AlibabaSellerVendorTradePurchaseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
