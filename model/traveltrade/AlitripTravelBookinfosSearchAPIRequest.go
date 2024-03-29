package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelBookinfosSearchAPIRequest 飞猪度假-订单预定信息列表搜索接口 API请求
// alitrip.travel.bookinfos.search
//
// 查询订单预定信息列表
type AlitripTravelBookinfosSearchAPIRequest struct {
	model.Params
	// 申请时间_结束，精确到分钟
	_applyTimeEnd string
	// 申请时间_开始，精确到分钟
	_applyTimeStart string
	// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
	_pageSize int64
	// 当前页
	_currentPage int64
}

// NewAlitripTravelBookinfosSearchRequest 初始化AlitripTravelBookinfosSearchAPIRequest对象
func NewAlitripTravelBookinfosSearchRequest() *AlitripTravelBookinfosSearchAPIRequest {
	return &AlitripTravelBookinfosSearchAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelBookinfosSearchAPIRequest) Reset() {
	r._applyTimeEnd = ""
	r._applyTimeStart = ""
	r._pageSize = 0
	r._currentPage = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelBookinfosSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.bookinfos.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelBookinfosSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelBookinfosSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyTimeEnd is ApplyTimeEnd Setter
// 申请时间_结束，精确到分钟
func (r *AlitripTravelBookinfosSearchAPIRequest) SetApplyTimeEnd(_applyTimeEnd string) error {
	r._applyTimeEnd = _applyTimeEnd
	r.Set("apply_time_end", _applyTimeEnd)
	return nil
}

// GetApplyTimeEnd ApplyTimeEnd Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetApplyTimeEnd() string {
	return r._applyTimeEnd
}

// SetApplyTimeStart is ApplyTimeStart Setter
// 申请时间_开始，精确到分钟
func (r *AlitripTravelBookinfosSearchAPIRequest) SetApplyTimeStart(_applyTimeStart string) error {
	r._applyTimeStart = _applyTimeStart
	r.Set("apply_time_start", _applyTimeStart)
	return nil
}

// GetApplyTimeStart ApplyTimeStart Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetApplyTimeStart() string {
	return r._applyTimeStart
}

// SetPageSize is PageSize Setter
// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
func (r *AlitripTravelBookinfosSearchAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetCurrentPage is CurrentPage Setter
// 当前页
func (r *AlitripTravelBookinfosSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

var poolAlitripTravelBookinfosSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelBookinfosSearchRequest()
	},
}

// GetAlitripTravelBookinfosSearchRequest 从 sync.Pool 获取 AlitripTravelBookinfosSearchAPIRequest
func GetAlitripTravelBookinfosSearchAPIRequest() *AlitripTravelBookinfosSearchAPIRequest {
	return poolAlitripTravelBookinfosSearchAPIRequest.Get().(*AlitripTravelBookinfosSearchAPIRequest)
}

// ReleaseAlitripTravelBookinfosSearchAPIRequest 将 AlitripTravelBookinfosSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelBookinfosSearchAPIRequest(v *AlitripTravelBookinfosSearchAPIRequest) {
	v.Reset()
	poolAlitripTravelBookinfosSearchAPIRequest.Put(v)
}
