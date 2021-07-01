package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallAscpOrdersSaleCreateAPIRequest
ASCP渠道中心销售单创建接口 API请求
tmall.ascp.orders.sale.create

ASCP渠道中心销售单创建接口 */
type TmallAscpOrdersSaleCreateAPIRequest struct {
	model.Params
	// 请求对象
	_channelOrderRequest *CreateChannelOrderRequest
}

// New
