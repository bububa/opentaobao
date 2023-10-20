package opentrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeActivityQueryAPIRequest 查询尖货活动信息 API请求
// taobao.opentrade.activity.query
//
// 尖货交易活动信息配置，查询尖货活动信息
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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpentradeActivityQueryAPIRequest) Reset() {
	r._endTime = ""
	r._activityName = ""
	r._pageSize = 0
	r._pageIndex = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeActivityQueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpentradeActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndTime is EndTime Setter
// 活动结束时间
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetActivityName is ActivityName Setter
// 活动名称
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetActivityName(_activityName string) error {
	r._activityName = _activityName
	r.Set("activity_name", _activityName)
	return nil
}

// GetActivityName ActivityName Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetActivityName() string {
	return r._activityName
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 分页序号
func (r *TaobaoOpentradeActivityQueryAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r TaobaoOpentradeActivityQueryAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

var poolTaobaoOpentradeActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpentradeActivityQueryRequest()
	},
}

// GetTaobaoOpentradeActivityQueryRequest 从 sync.Pool 获取 TaobaoOpentradeActivityQueryAPIRequest
func GetTaobaoOpentradeActivityQueryAPIRequest() *TaobaoOpentradeActivityQueryAPIRequest {
	return poolTaobaoOpentradeActivityQueryAPIRequest.Get().(*TaobaoOpentradeActivityQueryAPIRequest)
}

// ReleaseTaobaoOpentradeActivityQueryAPIRequest 将 TaobaoOpentradeActivityQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpentradeActivityQueryAPIRequest(v *TaobaoOpentradeActivityQueryAPIRequest) {
	v.Reset()
	poolTaobaoOpentradeActivityQueryAPIRequest.Put(v)
}
