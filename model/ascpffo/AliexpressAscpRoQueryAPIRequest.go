package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpRoQueryAPIRequest
AliExpress退供单查询API API请求
aliexpress.ascp.ro.query

AE仓发商家单个退供单查询接口 */
type AliexpressAscpRoQueryAPIRequest struct {
	model.Params
	// dto
	_returnOrderQuery *ReturnOrderQueryDto
}

// New
