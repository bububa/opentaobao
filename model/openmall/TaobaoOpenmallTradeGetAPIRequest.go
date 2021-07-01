package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeGetAPIRequest
查询订单详情 API请求
taobao.openmall.trade.get

查询订单详情 */
type TaobaoOpenmallTradeGetAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 淘宝订单号
	_tid int64
}

// New
