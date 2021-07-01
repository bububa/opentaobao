package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuCategoryAttrvalueGetAPIRequest
属性值获取 API请求
alibaba.icbu.category.attrvalue.get

属性值获取 */
type AlibabaIcbuCategoryAttrvalueGetAPIRequest struct {
	model.Params
	// 属性值request对象
	_attributeValueRequest *AttributeValueRequest
}

// New
