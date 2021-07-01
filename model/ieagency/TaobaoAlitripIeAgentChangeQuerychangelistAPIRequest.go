package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest
卖家查询改签列表 API请求
taobao.alitrip.ie.agent.change.querychangelist

提供B2B卖家查看改签列表服务 */
type TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest struct {
	model.Params
	// WAITING_CONFIRM(10, "卖家待确认"),CONFIRMED(20, "卖家已确认"),WAITING_ISSUE(30, "卖家待出票"),FROZEN_ORDER(40, "出票超时冻结"),ISSUE_SUCCESS(50, "出票成功"),CHECKING_FAILURE(60,"验真失败"),CHECKING_SUCCCESS(61,"验真成功"),REFUSED(70, "卖家已拒绝")
	_changeBizStatusEnum string
	// 改签申请单ID
	_changeOrderId int64
	// 申请原因类型 0－因乘客个人原因(自愿改签) ,1－因航班取消/延误(非自愿),
	_changeReasonType string
	// 1
	_endCreateDate string
	// 订单ID
	_orderId int64
	// 分页索引
	_pageIndex int64
	// 分页大小
	_pageSize int64
	// 排序
	_sortField int64
	// 1
	_startCreateDate string
}

// New
