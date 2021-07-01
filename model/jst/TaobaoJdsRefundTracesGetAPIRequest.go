package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJdsRefundTracesGetAPIRequest
获取单条退款跟踪详情 API请求
taobao.jds.refund.traces.get

获取聚石塔数据共享的交易全链路的退款信息 */
type TaobaoJdsRefundTracesGetAPIRequest struct {
	model.Params
	// 淘宝的退款编号
	_refundId int64
}

// New
