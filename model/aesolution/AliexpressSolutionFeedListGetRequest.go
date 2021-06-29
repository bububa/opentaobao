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
type AliexpressSolutionFeedListGetRequest struct {
    model.Params
    // current page
    currentPage   int64
    // feed type
    feedType   string
    // page size
    pageSize   int64
    // status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
    status   string
    // Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
    submittedTimeEnd   string
    // Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
    submittedTimeStart   string
}

// 初始化AliexpressSolutionFeedListGetRequest对象
func NewAliexpressSolutionFeedListGetRequest() *AliexpressSolutionFeedListGetRequest{
    return &AliexpressSolutionFeedListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedListGetRequest) GetApiMethodName() string {
    return "aliexpress.solution.feed.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentPage Setter
// current page
func (r *AliexpressSolutionFeedListGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

// CurrentPage Getter
func (r AliexpressSolutionFeedListGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}
// FeedType Setter
// feed type
func (r *AliexpressSolutionFeedListGetRequest) SetFeedType(feedType string) error {
    r.feedType = feedType
    r.Set("feed_type", feedType)
    return nil
}

// FeedType Getter
func (r AliexpressSolutionFeedListGetRequest) GetFeedType() string {
    return r.feedType
}
// PageSize Setter
// page size
func (r *AliexpressSolutionFeedListGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressSolutionFeedListGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// Status Setter
// status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
func (r *AliexpressSolutionFeedListGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AliexpressSolutionFeedListGetRequest) GetStatus() string {
    return r.status
}
// SubmittedTimeEnd Setter
// Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
func (r *AliexpressSolutionFeedListGetRequest) SetSubmittedTimeEnd(submittedTimeEnd string) error {
    r.submittedTimeEnd = submittedTimeEnd
    r.Set("submitted_time_end", submittedTimeEnd)
    return nil
}

// SubmittedTimeEnd Getter
func (r AliexpressSolutionFeedListGetRequest) GetSubmittedTimeEnd() string {
    return r.submittedTimeEnd
}
// SubmittedTimeStart Setter
// Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
func (r *AliexpressSolutionFeedListGetRequest) SetSubmittedTimeStart(submittedTimeStart string) error {
    r.submittedTimeStart = submittedTimeStart
    r.Set("submitted_time_start", submittedTimeStart)
    return nil
}

// SubmittedTimeStart Getter
func (r AliexpressSolutionFeedListGetRequest) GetSubmittedTimeStart() string {
    return r.submittedTimeStart
}
