package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeActivityQueryAPIRequest
查询尖货活动信息 API请求
taobao.opentrade.activity.query

尖货交易活动信息配置，查询尖货活动信息 */
type TaobaoOpentradeActivityQueryAPIRequest struct {
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

// NewTaobaoOpentradeActivityQueryRequest 初始化TaobaoOpentradeActivityQueryAPIRequest对象
func NewTaobaoOpentradeActivityQueryRequest() *TaobaoOpentradeActivityQueryAPIRequest {
	return &TaobaoOpentradeActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeActivityQueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeActivityQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EndTime Setter
// 活动结束时间
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is ActivityName Setter
// 活动名称
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// Get ActivityName Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetActivityName() string {
	return r._activityName
}

// Set is PageSize Setter
// 分页大小
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is PageIndex Setter
// 分页序号
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
