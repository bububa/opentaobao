package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuCategoryLevelAttrGetAPIRequest
层级属性的子属性获取 API请求
alibaba.icbu.category.level.attr.get

用于获取层级属性（车型库）的子属性和属性值 */
type AlibabaIcbuCategoryLevelAttrGetAPIRequest struct {
	model.Params
	// 属性值request对象
	_attributeValueRequest *LevelAttributeValueRequest
}

// New
