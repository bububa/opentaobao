package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpFroQueryAPIRequest
AliExpress销退单查询API API请求
aliexpress.ascp.fro.query

AE履约销退单查询接口 */
type AliexpressAscpFroQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentReverseOrderQuery *FulfillmentReverseOrderQueryDto
}

// New
