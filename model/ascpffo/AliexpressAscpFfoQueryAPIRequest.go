package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpFfoQueryAPIRequest
AliExpress发货单查询API API请求
aliexpress.ascp.ffo.query

AE 履约发货单分页查询接口 */
type AliexpressAscpFfoQueryAPIRequest struct {
	model.Params
	// dto
	_fulfillmentForwardOrderQuery *FulfillmentForwardOrderQueryDto
}

// New
