package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionFeedListGetAPIRequest
aliexpress.solution.feed.list.get API请求
aliexpress.solution.feed.list.get

API to query the feed list belonged to a seller */
type AliexpressSolutionFeedListGetAPIRequest struct {
	model.Params
	// current page
	_currentPage int64
	// feed type
	_feedType string
	// page size
	_pageSize int64
	// status of the job, currently there are 3 types: FINISH, PROCESSING, QUEUEING
	_status string
	// Search for feeds submitted before a specific time, format: yyyy-MM-dd hh:mm:ss. Timezone:America/Los_Angeles
	_submittedTimeEnd string
	// Search for feeds submitted after a specific time, format: yyyy-MM-dd hh:mm:ss  .Timezone:America/Los_Angeles
	_submittedTimeStart string
}

// New
