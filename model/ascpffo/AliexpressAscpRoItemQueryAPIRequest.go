package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpRoItemQueryAPIRequest
AliExpress退供单明细查询API API请求
aliexpress.ascp.ro.item.query

AE仓发 单个退供单明细查询 */
type AliexpressAscpRoItemQueryAPIRequest struct {
	model.Params
	// dto
	_returnOrderItemQuery *ReturnOrderItemQueryDto
}

// New
