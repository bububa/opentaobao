package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
层级属性的子属性获取 APIRequest
alibaba.icbu.category.level.attr.get

用于获取层级属性（车型库）的子属性和属性值
*/
type AlibabaIcbuCategoryLevelAttrGetRequest struct {
    model.Params

    // 属性值request对象
    attributeValueRequest   *LevelAttributeValueRequest 

}

func NewAlibabaIcbuCategoryLevelAttrGetRequest() *AlibabaIcbuCategoryLevelAttrGetRequest{
    return &AlibabaIcbuCategoryLevelAttrGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuCategoryLevelAttrGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.level.attr.get"
}

func (r AlibabaIcbuCategoryLevelAttrGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuCategoryLevelAttrGetRequest) SetAttributeValueRequest(attributeValueRequest *LevelAttributeValueRequest) error {
    r.attributeValueRequest = attributeValueRequest
    r.Set("attribute_value_request", attributeValueRequest)
    return nil
}

func (r AlibabaIcbuCategoryLevelAttrGetRequest) GetAttributeValueRequest() *LevelAttributeValueRequest {
    return r.attributeValueRequest
}

