package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductGroupGetAPIRequest
分组信息获取 API请求
alibaba.icbu.product.group.get

分组信息获取 */
type AlibabaIcbuProductGroupGetAPIRequest struct {
	model.Params
	// 分组ID，传-1获得所有一级分组
	_groupId int64
	// 补充信息
	_extraContext string
}

// New
