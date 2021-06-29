package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
属性值获取 API请求
alibaba.icbu.category.attrvalue.get

属性值获取
*/
type AlibabaIcbuCategoryAttrvalueGetRequest struct {
    model.Params
    // 属性值request对象
    attributeValueRequest   *AttributeValueRequest
}

// 初始化AlibabaIcbuCategoryAttrvalueGetRequest对象
func NewAlibabaIcbuCategoryAttrvalueGetRequest() *AlibabaIcbuCategoryAttrvalueGetRequest{
    return &AlibabaIcbuCategoryAttrvalueGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryAttrvalueGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.attrvalue.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryAttrvalueGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AttributeValueRequest Setter
// 属性值request对象
func (r *AlibabaIcbuCategoryAttrvalueGetRequest) SetAttributeValueRequest(attributeValueRequest *AttributeValueRequest) error {
    r.attributeValueRequest = attributeValueRequest
    r.Set("attribute_value_request", attributeValueRequest)
    return nil
}

// AttributeValueRequest Getter
func (r AlibabaIcbuCategoryAttrvalueGetRequest) GetAttributeValueRequest() *AttributeValueRequest {
    return r.attributeValueRequest
}
