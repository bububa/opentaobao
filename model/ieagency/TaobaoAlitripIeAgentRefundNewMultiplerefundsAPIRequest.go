package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIRequest
补退接口 API请求
taobao.alitrip.ie.agent.refund.new.multiplerefunds

1. 补退接口， 可以进行多次退款 */
type TaobaoAlitripIeAgentRefundNewMultiplerefundsAPIRequest struct {
	model.Params
	// 请求参数
	_paramRefundOrderMultipleRefundsRq *RefundOrderMultipleRefundsRq
}

// New
