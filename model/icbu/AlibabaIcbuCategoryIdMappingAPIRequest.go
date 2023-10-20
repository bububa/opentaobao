package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategoryidmappingAPIRequest 新旧属性的映射 API请求
// alibaba.icbu.category.id.mapping
//
// 商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
type AlibabaicbucategoryidmappingAPIRequest struct {
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

// NewAlibabaicbucategoryidmappingRequest 初始化AlibabaicbucategoryidmappingAPIRequest对象
func NewAlibabaicbucategoryidmappingRequest() *AlibabaicbucategoryidmappingAPIRequest {
	return &AlibabaicbucategoryidmappingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbucategoryidmappingAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.id.mapping"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbucategoryidmappingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbucategoryidmappingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id
func (r *AlibabaicbucategoryidmappingAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaicbucategoryidmappingAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetAttributeValueId is AttributeValueId Setter
// 属性值id
func (r *AlibabaicbucategoryidmappingAPIRequest) SetAttributeValueId(_attributeValueId int64) error {
	r._attributeValueId = _attributeValueId
	r.Set("attribute_value_id", _attributeValueId)
	return nil
}

// GetAttributeValueId AttributeValueId Getter
func (r AlibabaicbucategoryidmappingAPIRequest) GetAttributeValueId() int64 {
	return r._attributeValueId
}

// SetAttributeId is AttributeId Setter
// 属性id
func (r *AlibabaicbucategoryidmappingAPIRequest) SetAttributeId(_attributeId int64) error {
	r._attributeId = _attributeId
	r.Set("attribute_id", _attributeId)
	return nil
}

// GetAttributeId AttributeId Getter
func (r AlibabaicbucategoryidmappingAPIRequest) GetAttributeId() int64 {
	return r._attributeId
}

// SetConvertType is ConvertType Setter
// 转化类型, 1 = 转化类目id 2= 转化属性id 3= 转化属性值id
func (r *AlibabaicbucategoryidmappingAPIRequest) SetConvertType(_convertType int64) error {
	r._convertType = _convertType
	r.Set("convert_type", _convertType)
	return nil
}

// GetConvertType ConvertType Getter
func (r AlibabaicbucategoryidmappingAPIRequest) GetConvertType() int64 {
	return r._convertType
}
