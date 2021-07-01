package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuCategorySchemaLevelGetAPIRequest
(新)层级属性获取 API请求
alibaba.icbu.category.schema.level.get

将表单中层级属性的子属性返回 */
type AlibabaIcbuCategorySchemaLevelGetAPIRequest struct {
	model.Params
	// 类目id
	_catId int64
	// 返回的文案的语种，可以输入en_US或者zh
	_language string
	// 层级属性的当前层级属性
	_xml string
}

// New
