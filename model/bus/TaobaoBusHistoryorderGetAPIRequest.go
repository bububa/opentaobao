package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusHistoryorderGetAPIRequest
历史订单查询（对账） API请求
taobao.bus.historyorder.get

历史订单查询，对账接口 */
type TaobaoBusHistoryorderGetAPIRequest struct {
	model.Params
	// 开始时间 2017-04-23 13:33:43
	_fromDate string
	// 分页大小 不超过1w
	_pageSize int64
	// 结束时间 2017-04-23 13:33:43
	_toDate string
	// offline_ticket 线下自助机； online_ticket：线上售票； 空 代表查全部
	_type string
	// 第几页 从1开始
	_pageIndex int64
}

// New
