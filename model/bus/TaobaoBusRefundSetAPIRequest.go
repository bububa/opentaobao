package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusRefundSetAPIRequest
B2B退票申请接口 API请求
taobao.bus.refund.set

B2B业务支持退票 */
type TaobaoBusRefundSetAPIRequest struct {
	model.Params
	// 入参
	_param0 *B2BRefundOrderRq
}

// New
