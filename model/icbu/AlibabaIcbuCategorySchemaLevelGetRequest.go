package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)层级属性获取 API请求
alibaba.icbu.category.schema.level.get

将表单中层级属性的子属性返回
*/
type AlibabaIcbuCategorySchemaLevelGetRequest struct {
    model.Params
    // 类目id
    catId   int64
    // 返回的文案的语种，可以输入en_US或者zh
    language   string
    // 层级属性的当前层级属性
    xml   string
}

// 初始化AlibabaIcbuCategorySchemaLevelGetRequest对象
func NewAlibabaIcbuCategorySchemaLevelGetRequest() *AlibabaIcbuCategorySchemaLevelGetRequest{
    return &AlibabaIcbuCategorySchemaLevelGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetApiMethodName() string {
    return "alibaba.icbu.category.schema.level.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *AlibabaIcbuCategorySchemaLevelGetRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetCatId() int64 {
    return r.catId
}
// Language Setter
// 返回的文案的语种，可以输入en_US或者zh
func (r *AlibabaIcbuCategorySchemaLevelGetRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetLanguage() string {
    return r.language
}
// Xml Setter
// 层级属性的当前层级属性
func (r *AlibabaIcbuCategorySchemaLevelGetRequest) SetXml(xml string) error {
    r.xml = xml
    r.Set("xml", xml)
    return nil
}

// Xml Getter
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetXml() string {
    return r.xml
}
