package icbuseller

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerVendorOrderList 国际站服务市场订单列表接口
// alibaba.seller.vendor.order.list
//
// 返回服务商在服务市场的客户订单
func AlibabaSellerVendorOrderList(ctx context.Context, clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorOrderListAPIRequest, resp *icbuseller.AlibabaSellerVendorOrderListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
