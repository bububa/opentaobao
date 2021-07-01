package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeikeEserviceOrderGetAPIRequest
客服外包订单查询 API请求
taobao.weike.eservice.order.get

用于客服外包中服务商查询订单列表 */
type TaobaoWeikeEserviceOrderGetAPIRequest struct {
	model.Params
	// 订单服务开始日期
	_startDate string
	// 订单是否可以排班
	_schedulingState bool
	// 商家昵称
	_sellerNick string
	// 每页记录数（默认20，最大不超过20）
	_pageSize int64
	// 订单服务结束日期
	_endDate string
	// 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
	_orderIdList []int64
	// 页码（默认为1）
	_pageNum int64
}

// New
