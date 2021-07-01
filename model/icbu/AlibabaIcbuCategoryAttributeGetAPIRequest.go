package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuCategoryAttributeGetAPIRequest
类目属性获取 API请求
alibaba.icbu.category.attribute.get

根据类目ID获取系统定义的属性 */
type AlibabaIcbuCategoryAttributeGetAPIRequest struct {
	model.Params
	// 发布类目id
	_catId int64
}

// New
