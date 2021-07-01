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

// NewAlibabaTuikeWebUnionOrderQueryRequest 初始化AlibabaTuikeWebUnionOrderQueryAPIRequest对象
func NewAlibabaTuikeWebUnionOrderQueryRequest() *AlibabaTuikeWebUnionOrderQueryAPIRequest {
	return &AlibabaTuikeWebUnionOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTuikeWebUnionOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tuike.web.union.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTuikeWebUnionOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TimeType Setter
// 0 表示time为下单时间;1表示time为更新时间
func (r *AlibabaTuikeWebUnionOrderQueryAPIRequest) SetTimeType(_timeType int64) error {
	r._timeType = _timeType
	r.Set("time_type", _timeType)
	return nil
}

// Get TimeType Getter
func (r AlibabaTuikeWebUnionOrderQueryAPIRequest) GetTimeType() int64 {
	return r._timeType
}

// Set is StartTime Setter
// 13位时间戳
func (r *AlibabaTuikeWebUnionOrderQueryAPIRequest) SetStartTime(_startTime int64) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r AlibabaTuikeWebUnionOrderQueryAPIRequest) GetStartTime() int64 {
	return r._startTime
}

// Set is EndTime Setter
// 13位时间戳
func (r *AlibabaTuikeWebUnionOrderQueryAPIRequest) SetEndTime(_endTime int64) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r AlibabaTuikeWebUnionOrderQueryAPIRequest) GetEndTime() int64 {
	return r._endTime
}

// Set is PageOffset Setter
// 页码偏移
func (r *AlibabaTuikeWebUnionOrderQueryAPIRequest) SetPageOffset(_pageOffset int64) error {
	r._pageOffset = _pageOffset
	r.Set("page_offset", _pageOffset)
	return nil
}

// Get PageOffset Getter
func (r AlibabaTuikeWebUnionOrderQueryAPIRequest) GetPageOffset() int64 {
	return r._pageOffset
}

// Set is PageSize Setter
// 返回条数
func (r *AlibabaTuikeWebUnionOrderQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaTuikeWebUnionOrderQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
