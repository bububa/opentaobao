package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcPayAPIRequest
POS收银成功后订单同步 API请求
alibaba.mj.oc.pay

此API用于在银泰商场中，消费者在收银台收银/退款时， POS系统在收银或退款成功后，调用此接口进行订单同步 */
type AlibabaMjOcPayAPIRequest struct {
	model.Params
	// 订单数据
	_posOrder *PosOrderDto
}

// New
