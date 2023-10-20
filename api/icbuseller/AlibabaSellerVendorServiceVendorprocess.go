package icbuseller

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuseller"
)

// AlibabaSellerVendorServiceVendorprocess 服务商客户关联信息
// alibaba.seller.vendor.service.vendorprocess
//
// 服务商客户关联信息
func AlibabaSellerVendorServiceVendorprocess(clt *core.SDKClient, req *icbuseller.AlibabaSellerVendorServiceVendorprocessAPIRequest, resp *icbuseller.AlibabaSellerVendorServiceVendorprocessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
