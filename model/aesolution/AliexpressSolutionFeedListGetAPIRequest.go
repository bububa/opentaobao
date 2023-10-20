package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionfeedlistgetAPIRequest aliexpress.solution.feed.list.get API请求
// aliexpress.solution.feed.list.get
//
// API to query the feed list belonged to a seller
type AliexpresssolutionfeedlistgetAPIRequest struct {
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

// NewAliexpresssolutionfeedlistgetRequest 初始化AliexpresssolutionfeedlistgetAPIRequest对象
func NewAliexpresssolutionfeedlistgetRequest() *AliexpresssolutionfeedlistgetAPIRequest {
	return &AliexpresssolutionfeedlistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionfeedlistgetAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.feed.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionfeedlistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionfeedlistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeedType is FeedType Setter
// feed type
func (r *AliexpresssolutionfeedlistgetAPIRequest) SetFeedType(_feedType string) error {
	r._feedType = _feedType
	r.Set("feed_type", _feedType)
	return nil
}

// GetFeedType FeedType Getter
func (r AliexpresssolutionfeedlistgetAPIRequest) GetFeedType() string {
	return r._feedType
}

// SetStatus is Status Setter
// status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
func (r *AliexpresssolutionfeedlistgetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AliexpresssolutionfeedlistgetAPIRequest) GetStatus() string {
	return r._status
}

// SetSubmittedTimeEnd is SubmittedTimeEnd Setter
// Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
func (r *AliexpresssolutionfeedlistgetAPIRequest) SetSubmittedTimeEnd(_submittedTimeEnd string) error {
	r._submittedTimeEnd = _submittedTimeEnd
	r.Set("submitted_time_end", _submittedTimeEnd)
	return nil
}

// GetSubmittedTimeEnd SubmittedTimeEnd Getter
func (r AliexpresssolutionfeedlistgetAPIRequest) GetSubmittedTimeEnd() string {
	return r._submittedTimeEnd
}

// SetSubmittedTimeStart is SubmittedTimeStart Setter
// Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
func (r *AliexpresssolutionfeedlistgetAPIRequest) SetSubmittedTimeStart(_submittedTimeStart string) error {
	r._submittedTimeStart = _submittedTimeStart
	r.Set("submitted_time_start", _submittedTimeStart)
	return nil
}

// GetSubmittedTimeStart SubmittedTimeStart Getter
func (r AliexpresssolutionfeedlistgetAPIRequest) GetSubmittedTimeStart() string {
	return r._submittedTimeStart
}

// SetCurrentPage is CurrentPage Setter
// current page
func (r *AliexpresssolutionfeedlistgetAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AliexpresssolutionfeedlistgetAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// page size
func (r *AliexpresssolutionfeedlistgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpresssolutionfeedlistgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
