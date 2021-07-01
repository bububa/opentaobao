package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbOrderdetailDateGetAPIRequest
按照日期范围查询物流订单详情 API请求
taobao.wlb.orderdetail.date.get

外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情 */
type TaobaoWlbOrderdetailDateGetAPIRequest struct {
	model.Params
	// 创建时间起始
	_startTime string
	// 创建时间结束
	_endTime string
	// 分页大小
	_pageSize int64
	// 分页下标
	_pageNo int64
}

// New
