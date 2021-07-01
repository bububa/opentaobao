package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentShoppingPushAPIRequest
国际机票大卖家Shopping推送 API请求
taobao.alitrip.ie.agent.shopping.push

用于国际机票大卖家主动推送Shopping结果更新缓存报价。 */
type TaobaoAlitripIeAgentShoppingPushAPIRequest struct {
	model.Params
	// 政策推送结构体
	_param0 *ShoppingPushRq
}

// New
