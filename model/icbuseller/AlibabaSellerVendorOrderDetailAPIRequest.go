package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSellerVendorOrderDetailAPIRequest
国际站服务市场订单详情接口 API请求
alibaba.seller.vendor.order.detail

国际站服务市场订单列表接口 */
type AlibabaSellerVendorOrderDetailAPIRequest struct {
	model.Params
	// 订单编号
	_orderNo string
}

// New
