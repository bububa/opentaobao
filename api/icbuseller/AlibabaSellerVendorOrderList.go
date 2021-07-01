package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

/* AlibabaSellerVendorOrderList
国际站服务市场订单列表接口
alibaba.seller.vendor.order.list

返回服务商在服务市场的客户订单 */
func AlibabaSellerVendorOrderList(clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorOrderListAPIRequest, session string) (*icbuseller.AlibabaSellerVendorOrderListAPIResponse, error) {
	var resp icbuseller.AlibabaSellerVendorOrderListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
