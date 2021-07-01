package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest
快递下单 API请求
alibaba.onetouch.logistics.express.logistics.order.create

快递下单 */
type AlibabaOnetouchLogisticsExpressLogisticsOrderCreateAPIRequest struct {
	model.Params
	// 请求参数对象
	_paramnQuery *PlaceOrderDto
}

// New
