package tuike

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTuikeWebUnionOrderQueryAPIRequest
推客网盟合作抽佣订单查询接口 API请求
alibaba.tuike.web.union.order.query

推客网盟合作抽佣订单查询接口 */
type AlibabaTuikeWebUnionOrderQueryAPIRequest struct {
	model.Params
	// 0 表示time为下单时间;1表示time为更新时间
	_timeType int64
	// 13位时间戳
	_startTime int64
	// 13位时间戳
	_endTime int64
	// 页码偏移
	_pageOffset int64
	// 返回条数
	_pageSize int64
}

// New
