package aesolution

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionFeedListGetAPIRequest aliexpress.solution.feed.list.get API请求
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
type AliexpressSolutionFeedListGetAPIRequest struct {
	model.Params
	// feed type
	_feedType string
	// status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
	_status string
	// Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
	_submittedTimeEnd string
	// Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
	_submittedTimeStart string
	// current page
	_currentPage int64
	// page size
	_pageSize int64
}

// NewAliexpressSolutionFeedListGetRequest 初始化AliexpressSolutionFeedListGetAPIRequest对象
func NewAliexpressSolutionFeedListGetRequest() *AliexpressSolutionFeedListGetAPIRequest {
	return &AliexpressSolutionFeedListGetAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSolutionFeedListGetAPIRequest) Reset() {
	r._feedType = ""
	r._status = ""
	r._submittedTimeEnd = ""
	r._submittedTimeStart = ""
	r._currentPage = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionFeedListGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.feed.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionFeedListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSolutionFeedListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeedType is FeedType Setter
// feed type
func (r *AliexpressSolutionFeedListGetAPIRequest) SetFeedType(_feedType string) error {
	r._feedType = _feedType
	r.Set("feed_type", _feedType)
	return nil
}

// GetFeedType FeedType Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetFeedType() string {
	return r._feedType
}

// SetStatus is Status Setter
// status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
func (r *AliexpressSolutionFeedListGetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetStatus() string {
	return r._status
}

// SetSubmittedTimeEnd is SubmittedTimeEnd Setter
// Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
func (r *AliexpressSolutionFeedListGetAPIRequest) SetSubmittedTimeEnd(_submittedTimeEnd string) error {
	r._submittedTimeEnd = _submittedTimeEnd
	r.Set("submitted_time_end", _submittedTimeEnd)
	return nil
}

// GetSubmittedTimeEnd SubmittedTimeEnd Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetSubmittedTimeEnd() string {
	return r._submittedTimeEnd
}

// SetSubmittedTimeStart is SubmittedTimeStart Setter
// Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
func (r *AliexpressSolutionFeedListGetAPIRequest) SetSubmittedTimeStart(_submittedTimeStart string) error {
	r._submittedTimeStart = _submittedTimeStart
	r.Set("submitted_time_start", _submittedTimeStart)
	return nil
}

// GetSubmittedTimeStart SubmittedTimeStart Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetSubmittedTimeStart() string {
	return r._submittedTimeStart
}

// SetCurrentPage is CurrentPage Setter
// current page
func (r *AliexpressSolutionFeedListGetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// page size
func (r *AliexpressSolutionFeedListGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressSolutionFeedListGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAliexpressSolutionFeedListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSolutionFeedListGetRequest()
	},
}

// GetAliexpressSolutionFeedListGetRequest 从 sync.Pool 获取 AliexpressSolutionFeedListGetAPIRequest
func GetAliexpressSolutionFeedListGetAPIRequest() *AliexpressSolutionFeedListGetAPIRequest {
	return poolAliexpressSolutionFeedListGetAPIRequest.Get().(*AliexpressSolutionFeedListGetAPIRequest)
}

// ReleaseAliexpressSolutionFeedListGetAPIRequest 将 AliexpressSolutionFeedListGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressSolutionFeedListGetAPIRequest(v *AliexpressSolutionFeedListGetAPIRequest) {
	v.Reset()
	poolAliexpressSolutionFeedListGetAPIRequest.Put(v)
}
