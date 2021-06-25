package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
属性值获取 APIRequest
alibaba.icbu.category.attrvalue.get

属性值获取
*/
type AlibabaIcbuCategoryAttrvalueGetRequest struct {
    model.Params

    // 属性值request对象
    attributeValueRequest   *AttributeValueRequest 

}

func NewAlibabaIcbuCategoryAttrvalueGetRequest() *AlibabaIcbuCategoryAttrvalueGetRequest{
    return &AlibabaIcbuCategoryAttrvalueGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuCategoryAttrvalueGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.attrvalue.get"
}

func (r AlibabaIcbuCategoryAttrvalueGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuCategoryAttrvalueGetRequest) SetAttributeValueRequest(attributeValueRequest *AttributeValueRequest) error {
    r.attributeValueRequest = attributeValueRequest
    r.Set("attribute_value_request", attributeValueRequest)
    return nil
}

func (r AlibabaIcbuCategoryAttrvalueGetRequest) GetAttributeValueRequest() *AttributeValueRequest {
    return r.attributeValueRequest
}

