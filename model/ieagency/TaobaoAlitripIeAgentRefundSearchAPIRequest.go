package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundSearchAPIRequest
卖家查询退票申请 API请求
taobao.alitrip.ie.agent.refund.search

卖家查询退票申请 */
type TaobaoAlitripIeAgentRefundSearchAPIRequest struct {
	model.Params
	// 查询起始时间
	_createStartTime string
	// 查询结束时间
	_createEndTime string
	// WAIT(1,"待处理"), AGREED(2, "已同意"),REFUSE(3, "已拒绝"),PROCESS(6, "已受理"), SUCCESS(7, "已退款");
	_refundStatus int64
	// 从1开始
	_pageIndex int64
	// 每页大小
	_pageSize int64
	// 代理商id
	_agentId int64
}

// New
