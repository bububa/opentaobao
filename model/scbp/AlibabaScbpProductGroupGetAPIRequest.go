package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpProductGroupGetAPIRequest
查询指定产品分组的下一层子分组 API请求
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组 */
type AlibabaScbpProductGroupGetAPIRequest struct {
	model.Params
	// 产品分组标识，null表示查询第一层分组
	_groupId string
}

// New
