package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpFfoItemQueryAPIRequest
AliExpress发货单明细分页查询API API请求
aliexpress.ascp.ffo.item.query

AE履约发货单明细分页查询 */
type AliexpressAscpFfoItemQueryAPIRequest struct {
	model.Params
	// DTO
	_fulfillmentForwardOrderItemQuery *FulfillmentForwardOrderItemQueryDto
}

// New
