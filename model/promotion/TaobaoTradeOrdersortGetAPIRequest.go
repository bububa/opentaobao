package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeOrdersortGetAPIRequest
获取前N有礼活动的开奖订单列表 API请求
taobao.trade.ordersort.get

获取前N有礼活动的开奖订单列表 */
type TaobaoTradeOrdersortGetAPIRequest struct {
	model.Params
	// 活动ID
	_activityId int64
	// 页码
	_pageNo int64
	// 一页记录数, 必须写死500
	_pageSize int64
}

// New
