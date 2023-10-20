package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerVendorOrderDetail 国际站服务市场订单详情接口
// alibaba.seller.vendor.order.detail
//
// 国际站服务市场订单列表接口
func AlibabaSellerVendorOrderDetail(clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorOrderDetailAPIRequest, session string) (*icbuseller.AlibabaSellerVendorOrderDetailAPIResponse, error) {
	var resp icbuseller.AlibabaSellerVendorOrderDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
