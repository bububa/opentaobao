package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelTradeQueryAPIRequest
飞猪度假-订单详情查询接口 API请求
alitrip.travel.trade.query

飞猪度假订单详情查询接口 */
type AlitripTravelTradeQueryAPIRequest struct {
	model.Params
	// 主订单id
	_orderId int64
}

// New
