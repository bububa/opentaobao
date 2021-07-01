package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeCloseAPIRequest
关闭订单 API请求
taobao.openmall.trade.close

关闭订单 */
type TaobaoOpenmallTradeCloseAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 关单原因
	_reason string
	// 淘宝订单号
	_tid int64
}

// New
