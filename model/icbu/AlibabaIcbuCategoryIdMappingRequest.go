package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新旧属性的映射 API请求
alibaba.icbu.category.id.mapping

商品发布接口升级，需要传入新的类目。这个接口 根据旧的类目id，获取新的类目id
*/
type AlibabaIcbuCategoryIdMappingRequest struct {
    model.Params
    // 发布类目id
    _catId   int64
    // 属性值id
    _attributeValueId   int64
    // 属性id
    _attributeId   int64
    // 转化类型, 1 = 转化类目id 2= 转化属性id 3= 转化属性值id
    _convertType   int64
}

// 初始化AlibabaIcbuCategoryIdMappingRequest对象
func NewAlibabaIcbuCategoryIdMappingRequest() *AlibabaIcbuCategoryIdMappingRequest{
    return &AlibabaIcbuCategoryIdMappingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryIdMappingRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.id.mapping"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryIdMappingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 发布类目id
func (r *AlibabaIcbuCategoryIdMappingRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategoryIdMappingRequest) GetCatId() int64 {
    return r._catId
}
// AttributeValueId Setter
// 属性值id
func (r *AlibabaIcbuCategoryIdMappingRequest) SetAttributeValueId(_attributeValueId int64) error {
    r._attributeValueId = _attributeValueId
    r.Set("attribute_value_id", _attributeValueId)
    return nil
}

// AttributeValueId Getter
func (r AlibabaIcbuCategoryIdMappingRequest) GetAttributeValueId() int64 {
    return r._attributeValueId
}
// AttributeId Setter
// 属性id
func (r *AlibabaIcbuCategoryIdMappingRequest) SetAttributeId(_attributeId int64) error {
    r._attributeId = _attributeId
    r.Set("attribute_id", _attributeId)
    return nil
}

// AttributeId Getter
func (r AlibabaIcbuCategoryIdMappingRequest) GetAttributeId() int64 {
    return r._attributeId
}
// ConvertType Setter
// 转化类型, 1 = 转化类目id 2= 转化属性id 3= 转化属性值id
func (r *AlibabaIcbuCategoryIdMappingRequest) SetConvertType(_convertType int64) error {
    r._convertType = _convertType
    r.Set("convert_type", _convertType)
    return nil
}

// ConvertType Getter
func (r AlibabaIcbuCategoryIdMappingRequest) GetConvertType() int64 {
    return r._convertType
}
