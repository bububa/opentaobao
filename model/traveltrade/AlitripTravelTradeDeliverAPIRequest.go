package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelTradeDeliverAPIRequest
飞猪度假-订单发货接口 API请求
alitrip.travel.trade.deliver

航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货 */
type AlitripTravelTradeDeliverAPIRequest struct {
	model.Params
	// 子订单id
	_subOrderId int64
}

// New
