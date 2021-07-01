package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
使用schema模板进行商品发布 API请求
alitrip.item.schema.add

飞猪度假商品使用schema模板进行商品发布。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemSchemaAddAPIRequest struct {
    model.Params
    // 类目id
    _catId   int64
    // 商品数据
    _schemaXmlFields   string
}

// 初始化AlitripItemSchemaAddAPIRequest对象
func NewAlitripItemSchemaAddRequest() *AlitripItemSchemaAddAPIRequest{
    return &AlitripItemSchemaAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripItemSchemaAddAPIRequest) GetApiMethodName() string {
    return "alitrip.item.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripItemSchemaAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *AlitripItemSchemaAddAPIRequest) SetCatId(_catId int64) error {
    r._catId = _catId
    r.Set("cat_id", _catId)
    return nil
}

// CatId Getter
func (r AlitripItemSchemaAddAPIRequest) GetCatId() int64 {
    return r._catId
}
// SchemaXmlFields Setter
// 商品数据
func (r *AlitripItemSchemaAddAPIRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
    r._schemaXmlFields = _schemaXmlFields
    r.Set("schema_xml_fields", _schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r AlitripItemSchemaAddAPIRequest) GetSchemaXmlFields() string {
    return r._schemaXmlFields
}
