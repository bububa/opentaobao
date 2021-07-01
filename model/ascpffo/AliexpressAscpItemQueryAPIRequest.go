package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAscpItemQueryAPIRequest
AliExpress货品查询查询API API请求
aliexpress.ascp.item.query

AE货品查询API */
type AliexpressAscpItemQueryAPIRequest struct {
	model.Params
	// DTO
	_scItemQuery *ScItemQueryDto
}

// New
