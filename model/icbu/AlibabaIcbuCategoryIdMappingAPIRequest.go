package icbu

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) Reset() {
	r._catId = 0
	r._attributeValueId = 0
	r._attributeId = 0
	r._convertType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.id.mapping"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetAttributeValueId is AttributeValueId Setter
// 属性值id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetAttributeValueId(_attributeValueId int64) error {
	r._attributeValueId = _attributeValueId
	r.Set("attribute_value_id", _attributeValueId)
	return nil
}

// GetAttributeValueId AttributeValueId Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetAttributeValueId() int64 {
	return r._attributeValueId
}

// SetAttributeId is AttributeId Setter
// 属性id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetAttributeId(_attributeId int64) error {
	r._attributeId = _attributeId
	r.Set("attribute_id", _attributeId)
	return nil
}

// GetAttributeId AttributeId Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetAttributeId() int64 {
	return r._attributeId
}

// SetConvertType is ConvertType Setter
// 转化类型, 1 = 转化类目id 2= 转化属性id 3= 转化属性值id
func (r *AlibabaIcbuCategoryIdMappingAPIRequest) SetConvertType(_convertType int64) error {
	r._convertType = _convertType
	r.Set("convert_type", _convertType)
	return nil
}

// GetConvertType ConvertType Getter
func (r AlibabaIcbuCategoryIdMappingAPIRequest) GetConvertType() int64 {
	return r._convertType
}

var poolAlibabaIcbuCategoryIdMappingAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuCategoryIdMappingRequest()
	},
}

// GetAlibabaIcbuCategoryIdMappingRequest 从 sync.Pool 获取 AlibabaIcbuCategoryIdMappingAPIRequest
func GetAlibabaIcbuCategoryIdMappingAPIRequest() *AlibabaIcbuCategoryIdMappingAPIRequest {
	return poolAlibabaIcbuCategoryIdMappingAPIRequest.Get().(*AlibabaIcbuCategoryIdMappingAPIRequest)
}

// ReleaseAlibabaIcbuCategoryIdMappingAPIRequest 将 AlibabaIcbuCategoryIdMappingAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuCategoryIdMappingAPIRequest(v *AlibabaIcbuCategoryIdMappingAPIRequest) {
	v.Reset()
	poolAlibabaIcbuCategoryIdMappingAPIRequest.Put(v)
}
