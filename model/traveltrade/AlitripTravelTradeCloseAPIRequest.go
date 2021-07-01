package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelTradeCloseAPIRequest
飞猪度假-订单关闭接口（快速退款） API请求
alitrip.travel.trade.close

卖家关单（快速退款接口），不支持二次预约商品的订单 */
type AlitripTravelTradeCloseAPIRequest struct {
	model.Params
	// 交易关闭原因
	_closeReason string
	// 子订单编号
	_subOrderId int64
	// 订单关闭原因描述
	_reasonDesc string
}

// New
