package aesolution

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
aliexpress.solution.feed.list.get API请求
aliexpress.solution.feed.list.get

API to query the feed list belonged to a seller
*/
type AliexpressSolutionFeedListGetAPIRequest struct {
    model.Params
    // current page
    _currentPage   int64
    // feed type
    _feedType   string
    // page size
    _pageSize   int64
    // status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
    _status   string
    // Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
    _submittedTimeEnd   string
    // Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
    _submittedTimeStart   string
}

// 初始化AliexpressSolutionFeedListGetAPIRequest对象
func NewAliexpressSolutionFeedListGetRequest() *AliexpressSolutionFeedListGetAPIRequest{
    return &AliexpressSolutionFeedListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedListGetAPIRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentPage Setter
// current page
func (r *AliexpressSolutionFeedListGetAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// FeedType Setter
// feed type
func (r *AliexpressSolutionFeedListGetAPIRequest) SetFeedType(_feedType string) error {
    r._feedType = _feedType
    r.Set("feed_type", _feedType)
    return nil
}

// FeedType Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetFeedType() string {
    return r._feedType
}
// PageSize Setter
// page size
func (r *AliexpressSolutionFeedListGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Status Setter
// status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
func (r *AliexpressSolutionFeedListGetAPIRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetStatus() string {
    return r._status
}
// SubmittedTimeEnd Setter
// Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
func (r *AliexpressSolutionFeedListGetAPIRequest) SetSubmittedTimeEnd(_submittedTimeEnd string) error {
    r._submittedTimeEnd = _submittedTimeEnd
    r.Set("submitted_time_end", _submittedTimeEnd)
    return nil
}

// SubmittedTimeEnd Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetSubmittedTimeEnd() string {
    return r._submittedTimeEnd
}
// SubmittedTimeStart Setter
// Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
func (r *AliexpressSolutionFeedListGetAPIRequest) SetSubmittedTimeStart(_submittedTimeStart string) error {
    r._submittedTimeStart = _submittedTimeStart
    r.Set("submitted_time_start", _submittedTimeStart)
    return nil
}

// SubmittedTimeStart Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetSubmittedTimeStart() string {
    return r._submittedTimeStart
}
