package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailTradeCloseticketAPIRequest
出票失败关单接口 API请求
alitrip.rail.trade.closeticket

出票成功回调接口 */
type AlitripRailTradeCloseticketAPIRequest struct {
	model.Params
	// 平台订单号
	_tpOrderId int64
	// 代理商订单号
	_agentId int64
	// 出票失败原因
	_errorMsg string
	// 出票失败码
	_errorCode string
}

// New
