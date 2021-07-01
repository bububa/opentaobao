package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpFroItemQueryAPIRequest
AliExpress销退单明细查询API API请求
aliexpress.ascp.fro.item.query

AE履约销退单明细查询API */
type AliexpressAscpFroItemQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentReverseOrderItemQuery *FulfillmentReverseOrderItemQueryDto
}

// New
