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
type AlitripItemSchemaAddRequest struct {
    model.Params
    // 类目id
    catId   int64
    // 商品数据
    schemaXmlFields   string
}

// 初始化AlitripItemSchemaAddRequest对象
func NewAlitripItemSchemaAddRequest() *AlitripItemSchemaAddRequest{
    return &AlitripItemSchemaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripItemSchemaAddRequest) GetApiMethodName() string {
    return "alitrip.item.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripItemSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CatId Setter
// 类目id
func (r *AlitripItemSchemaAddRequest) SetCatId(catId int64) error {
    r.catId = catId
    r.Set("cat_id", catId)
    return nil
}

// CatId Getter
func (r AlitripItemSchemaAddRequest) GetCatId() int64 {
    return r.catId
}
// SchemaXmlFields Setter
// 商品数据
func (r *AlitripItemSchemaAddRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r AlitripItemSchemaAddRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}
