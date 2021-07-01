package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleApprizeOrderFulfillmentAPIRequest
鉴定担保资金订单履约 API请求
alibaba.idle.apprize.order.fulfillment

服务商针对自己的服务订单进行履约 */
type AlibabaIdleApprizeOrderFulfillmentAPIRequest struct {
	model.Params
	// deal：服务商收取费用、refund：退款给付款方
	_action string
	// 天猫服务工单Id
	_workCardId int64
}

// New
