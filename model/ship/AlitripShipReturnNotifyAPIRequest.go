package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripShipReturnNotifyAPIRequest
船票退票退款回填接口 API请求
alitrip.ship.return.notify

此接口为接入商调用飞猪接口回填退票状态，飞猪平台给用户进行退票退款。飞猪平台保证数据幂等。 */
type AlitripShipReturnNotifyAPIRequest struct {
	model.Params
	// 退票请求入参
	_confirmRefundRQ *ShipAgentConfirmRefundRq
}

// New
