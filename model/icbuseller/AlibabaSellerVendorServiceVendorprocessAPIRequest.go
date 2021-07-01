package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSellerVendorServiceVendorprocessAPIRequest
服务商客户关联信息 API请求
alibaba.seller.vendor.service.vendorprocess

服务商客户关联信息 */
type AlibabaSellerVendorServiceVendorprocessAPIRequest struct {
	model.Params
	// order_num
	_orderNum string
}

// New
