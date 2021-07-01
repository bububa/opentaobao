package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionFeedInvalidateAPIRequest
aliexpress.solution.feed.invalidate API请求
aliexpress.solution.feed.invalidate

Api for invalidating specific feeds based on job Ids. Please use aliexpress.solution.feed.list.get to determine which job Ids needs to be sent for invalidation. */
type AliexpressSolutionFeedInvalidateAPIRequest struct {
	model.Params
	// job id separated by ;  No more than 100 job Ids could be passed in one request.
	_jobIdList string
}

// New
