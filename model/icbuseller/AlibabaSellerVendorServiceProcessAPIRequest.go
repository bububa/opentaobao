package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSellerVendorServiceProcessAPIRequest
服务商客户关联信息 API请求
alibaba.seller.vendor.service.process

服务商客户关联信息 */
type AlibabaSellerVendorServiceProcessAPIRequest struct {
	model.Params
	// order_num
	_orderNum string
}

// New
