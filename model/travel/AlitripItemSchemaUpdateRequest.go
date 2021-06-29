package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
使用schema进行商品编辑 API请求
alitrip.item.schema.update

飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemSchemaUpdateRequest struct {
    model.Params
    // 商品id
    _itemId   int64
    // 商品数据
    _schemaXmlFields   string
}

// 初始化AlitripItemSchemaUpdateRequest对象
func NewAlitripItemSchemaUpdateRequest() *AlitripItemSchemaUpdateRequest{
    return &AlitripItemSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripItemSchemaUpdateRequest) GetApiMethodName() string {
    return "alitrip.item.schema.update"
}

// IRequest interface 方法, 获取API参数
func (r AlitripItemSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *AlitripItemSchemaUpdateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlitripItemSchemaUpdateRequest) GetItemId() int64 {
    return r._itemId
}
// SchemaXmlFields Setter
// 商品数据
func (r *AlitripItemSchemaUpdateRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
    r._schemaXmlFields = _schemaXmlFields
    r.Set("schema_xml_fields", _schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r AlitripItemSchemaUpdateRequest) GetSchemaXmlFields() string {
    return r._schemaXmlFields
}
