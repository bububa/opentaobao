package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpProductListAPIRequest
查询P4P产品 API请求
alibaba.scbp.product.list

查询P4P产品 */
type AlibabaScbpProductListAPIRequest struct {
	model.Params
	// 产品分组标识
	_groupId string
	// 产品分页查询，每页个数，最大值20
	_perPageSize int64
	// 第几页
	_toPage int64
}

// New
