package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpSalecategoryQueryAPIRequest
货品品类查询 API请求
alibaba.ascp.salecategory.query

根据货品ID查询对应销售品类ID */
type AlibabaAscpSalecategoryQueryAPIRequest struct {
	model.Params
	// 货品ID
	_itemId []int64
}

// New
