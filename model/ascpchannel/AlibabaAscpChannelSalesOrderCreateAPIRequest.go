package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpChannelSalesOrderCreateAPIRequest
供应链渠道销售单创建接口 API请求
alibaba.ascp.channel.sales.order.create

阿里巴巴供应链渠道销售订单创建接口 */
type AlibabaAscpChannelSalesOrderCreateAPIRequest struct {
	model.Params
	// 请求参数
	_createOrderRequest *ExternalCreateSalesOrderRequest
}

// New
