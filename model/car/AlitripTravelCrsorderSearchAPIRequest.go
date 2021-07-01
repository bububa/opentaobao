package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelCrsorderSearchAPIRequest
CRS接送机订单列表搜索 API请求
alitrip.travel.crsorder.search

提供给CRS商家搜索订单列表，仅返回订单号列表 */
type AlitripTravelCrsorderSearchAPIRequest struct {
	model.Params
	// 订单状态，10-待派单，20-待用车，30-已取消，40-待处理退款申请，60-已关闭，70-已完成
	_crsOrderStatus int64
	// 用车时间-起始
	_beginCarUseTime string
	// 页大小，默认20
	_pageSize int64
	// 用车时间-终止
	_endCarUseTime string
	// 当前页，默认值1
	_currentPage int64
	// 支付时间-终止
	_endPayTime string
	// 支付时间-起始
	_beginPayTime string
	// 取消时间-起始
	_beginCancelTime string
	// 取消时间-终止
	_endCancelTime string
}

// New
