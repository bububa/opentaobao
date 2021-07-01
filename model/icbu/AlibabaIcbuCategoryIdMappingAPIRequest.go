package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuCategoryIdMappingAPIRequest
新旧属性的映射 API请求
alibaba.icbu.category.id.mapping

商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id */
type AlibabaIcbuCategoryIdMappingAPIRequest struct {
	model.Params
	// 发布类目id
	_catId int64
	// 属性值id
	_attributeValueId int64
	// 属性id
	_attributeId int64
	// 转化类型, 1 = 转化类目id 2= 转化属性id 3= 转化属性值id
	_convertType int64
}

// New
