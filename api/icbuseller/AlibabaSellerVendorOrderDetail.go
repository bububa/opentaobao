package icbuseller

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerVendorOrderDetail 国际站服务市场订单详情接口
// alibaba.seller.vendor.order.detail
//
// 国际站服务市场订单列表接口
func AlibabaSellerVendorOrderDetail(ctx context.Context, clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorOrderDetailAPIRequest, resp *icbuseller.AlibabaSellerVendorOrderDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
