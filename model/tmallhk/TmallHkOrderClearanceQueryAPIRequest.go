package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallHkOrderClearanceQueryAPIRequest
天猫国际订单清关信息 API请求
tmall.hk.order.clearance.query

天猫国际订单清关信息查询 */
type TmallHkOrderClearanceQueryAPIRequest struct {
	model.Params
	// 交易主订单号
	_bizOrderId int64
	// 调用方业务身份(由国际侧配置提供给调用方)
	_businessSymbol string
}

// New
