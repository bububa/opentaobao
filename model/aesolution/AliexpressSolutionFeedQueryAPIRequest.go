package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionFeedQueryAPIRequest
aliexpress.solution.feed.query API请求
aliexpress.solution.feed.query

API for query the execution result of feed. */
type AliexpressSolutionFeedQueryAPIRequest struct {
	model.Params
	// job id
	_jobId int64
}

// New
