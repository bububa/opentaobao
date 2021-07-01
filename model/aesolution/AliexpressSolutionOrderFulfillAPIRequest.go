package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionOrderFulfillAPIRequest
fulfill order API请求
aliexpress.solution.order.fulfill

fulfill order for seller */
type AliexpressSolutionOrderFulfillAPIRequest struct {
	model.Params
	// Actual logistics service selected by the user (logistics service key: This interface obtains the currently supportable logistics according to all the supportable logistics services listed by api.listLogisticsService. Please visit the forum link http://bbs.seller.aliexpress.com/bbs/read.php?tid=266120&page=1&toread=1#tpc for the detailed list of logistics services supported by the platform.)
	_serviceName string
	// When serviceName=other, fill in the corresponding tracking website.
	_trackingWebsite string
	// order ID for delivery by the user
	_outRef string
	// Status including: all shipments (all), part of the delivery (part)
	_sendType string
	// Remarks (only in English, and the length is limited to 512 characters)
	_description string
	// logistics number
	_logisticsNo string
}

// New
