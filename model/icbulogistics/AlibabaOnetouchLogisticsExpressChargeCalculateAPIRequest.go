package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest
计算快递运费&下单参数校验 API请求
alibaba.onetouch.logistics.express.charge.calculate

计算快递运费、下单参数校验 */
type AlibabaOnetouchLogisticsExpressChargeCalculateAPIRequest struct {
	model.Params
	// 请求参数对象
	_paramnQuery *PlaceOrderDto
}

// New
