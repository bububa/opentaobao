package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductGroupAddAPIRequest
增加商品分组 API请求
alibaba.icbu.product.group.add

增加商品分组 */
type AlibabaIcbuProductGroupAddAPIRequest struct {
	model.Params
	// 分组名称
	_groupName string
	// 上级分组ID，如果建立顶级分组设为-1
	_parentId int64
	// 补充信息，如isv id
	_extraContext string
}

// New
