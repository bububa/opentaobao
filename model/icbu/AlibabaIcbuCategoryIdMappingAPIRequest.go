package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryIdMappingAPIRequest 新旧属性的映射 API请求
// alibaba.icbu.category.id.mapping
//
// 商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
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

// NewAlibabaIcbuCategoryIdMappingRequest 初始化AlibabaIcbuCategoryIdMappingAPIRequest对象
func NewAlibabaIcbuCategoryIdMappingRequest() *AlibabaIcbuCategoryIdMappingAPIRequest {
	return &AlibabaIcbuCategoryIdMappingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.id.mapping"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CatId Setter
// 发布类目id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// Get CatId Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetCatId() int64 {
	return r._catId
}

// Set is AttributeValueId Setter
// 属性值id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetAttributeValueId(_attributeValueId int64) error {
	r._attributeValueId = _attributeValueId
	r.Set("attribute_value_id", _attributeValueId)
	return nil
}

// Get AttributeValueId Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetAttributeValueId() int64 {
	return r._attributeValueId
}

// Set is AttributeId Setter
// 属性id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetAttributeId(_attributeId int64) error {
	r._attributeId = _attributeId
	r.Set("attribute_id", _attributeId)
	return nil
}

// Get AttributeId Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetAttributeId() int64 {
	return r._attributeId
}

// Set is ConvertType Setter
// 转化类型, 1 = 转化类目id 2= 转化属性id 3= 转化属性值id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetConvertType(_convertType int64) error {
	r._convertType = _convertType
	r.Set("convert_type", _convertType)
	return nil
}

// Get ConvertType Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetConvertType() int64 {
	return r._convertType
}
