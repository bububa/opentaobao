package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSellerVendorOrderListAPIRequest
国际站服务市场订单列表接口 API请求
alibaba.seller.vendor.order.list

返回服务商在服务市场的客户订单 */
type AlibabaSellerVendorOrderListAPIRequest struct {
	model.Params
	// 查询参数
	_queryTradeDto *QueryTradeDto
}

// New
