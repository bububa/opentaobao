package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单预定信息列表搜索接口 API请求
alitrip.travel.bookinfos.search

查询订单预定信息列表
*/
type AlitripTravelBookinfosSearchAPIRequest struct {
    model.Params
    // 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
    _pageSize   int64
    // 当前页
    _currentPage   int64
    // 申请时间_结束，精确到分钟
    _applyTimeEnd   string
    // 申请时间_开始，精确到分钟
    _applyTimeStart   string
}

// 初始化AlitripTravelBookinfosSearchAPIRequest对象
func NewAlitripTravelBookinfosSearchRequest() *AlitripTravelBookinfosSearchAPIRequest{
    return &AlitripTravelBookinfosSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelBookinfosSearchAPIRequest) GetApiMethodName() string {
    return "alitrip.travel.bookinfos.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelBookinfosSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
func (r *AlitripTravelBookinfosSearchAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// CurrentPage Setter
// 当前页
func (r *AlitripTravelBookinfosSearchAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// ApplyTimeEnd Setter
// 申请时间_结束，精确到分钟
func (r *AlitripTravelBookinfosSearchAPIRequest) SetApplyTimeEnd(_applyTimeEnd string) error {
    r._applyTimeEnd = _applyTimeEnd
    r.Set("apply_time_end", _applyTimeEnd)
    return nil
}

// ApplyTimeEnd Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetApplyTimeEnd() string {
    return r._applyTimeEnd
}
// ApplyTimeStart Setter
// 申请时间_开始，精确到分钟
func (r *AlitripTravelBookinfosSearchAPIRequest) SetApplyTimeStart(_applyTimeStart string) error {
    r._applyTimeStart = _applyTimeStart
    r.Set("apply_time_start", _applyTimeStart)
    return nil
}

// ApplyTimeStart Getter
func (r AlitripTravelBookinfosSearchAPIRequest) GetApplyTimeStart() string {
    return r._applyTimeStart
}
