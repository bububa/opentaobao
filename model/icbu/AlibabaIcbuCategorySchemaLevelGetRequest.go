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
    _catId   int64
    // 返回的文案的语种，可以输入en_US或者zh
    _language   string
    // 层级属性的当前层级属性
    _xml   string
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
func (r *AlibabaIcbuCategorySchemaLevelGetRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetCatId() int64 {
    return r._catId
}
// Language Setter
// 返回的文案的语种，可以输入en_US或者zh
func (r *AlibabaIcbuCategorySchemaLevelGetRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetLanguage() string {
    return r._language
}
// Xml Setter
// 层级属性的当前层级属性
func (r *AlibabaIcbuCategorySchemaLevelGetRequest) SetXml(_xml string) error {
    r._xml = _xml
    r.Set("xml", _xml)
    return nil
}

// Xml Getter
func (r AlibabaIcbuCategorySchemaLevelGetRequest) GetXml() string {
    return r._xml
}
