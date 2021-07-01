package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionFeedSubmitAPIRequest
aliexpress.solution.feed.submit API请求
aliexpress.solution.feed.submit

API for merchants to submit feed data. Please note for each seller, the recommended number of feeds submitted for each operation_type every 24 hours should be lee than 150, otherwise significant delay might be encountered for processing the feed. */
type AliexpressSolutionFeedSubmitAPIRequest struct {
	model.Params
	// Currently support 4 types of feeds:PRODUCT_CREATE,PRODUCT_FULL_UPDATE,PRODUCT_STOCKS_UPDATE,PRODUCT_PRICES_UPDATE
	_operationType string
	// item list, maximum size: 2000.
	_itemList []SingleItemRequestDto
}

// New
