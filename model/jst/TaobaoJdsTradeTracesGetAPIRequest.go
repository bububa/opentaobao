package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJdsTradeTracesGetAPIRequest
获取单条订单跟踪详情 API请求
taobao.jds.trade.traces.get

获取聚石塔数据共享的交易全链路信息 */
type TaobaoJdsTradeTracesGetAPIRequest struct {
	model.Params
	// 淘宝的订单tid
	_tid int64
}

// New
