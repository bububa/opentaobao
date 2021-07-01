package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTripvpAgentOrderSearchAPIRequest
【国际机票】查询辅营订单列表 API请求
alitrip.tripvp.agent.order.search

【国际机票】查询辅营订单列表 */
type AlitripTripvpAgentOrderSearchAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 辅营创建开始时间
	_beginTime string
	// 当前页码
	_currentPage int64
	// 辅营创建结束时间
	_endTime string
	// 订单状态，1-待支付 2-支付成功 3-	辅营出货成功 4-订单取消
	_orderStatus int64
	// 分页行数
	_pageSize int64
}

// New
