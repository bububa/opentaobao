package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopentradeactivityqueryAPIRequest 查询尖货活动信息 API请求
// taobao.opentrade.activity.query
//
// 尖货交易活动信息配置，查询尖货活动信息
type TaobaoopentradeactivityqueryAPIRequest struct {
	model.Params
	// 活动结束时间
	_endTime string
	// 活动名称
	_activityName string
	// 分页大小
	_pageSize int64
	// 分页序号
	_pageIndex int64
}

// NewTaobaoopentradeactivityqueryRequest 初始化TaobaoopentradeactivityqueryAPIRequest对象
func NewTaobaoopentradeactivityqueryRequest() *TaobaoopentradeactivityqueryAPIRequest {
	return &TaobaoopentradeactivityqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopentradeactivityqueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopentradeactivityqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopentradeactivityqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndTime is EndTime Setter
// 活动结束时间
func (r *TaobaoopentradeactivityqueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoopentradeactivityqueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *TaobaoopentradeactivityqueryAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r TaobaoopentradeactivityqueryAPIRequest) GetActivityName() string {
	return r._activityName
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoopentradeactivityqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoopentradeactivityqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 分页序号
func (r *TaobaoopentradeactivityqueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoopentradeactivityqueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
