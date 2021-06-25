package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
使用schema进行商品编辑 APIRequest
alitrip.item.schema.update

飞猪度假商品使用schema进行商品编辑。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
*/
type AlitripItemSchemaUpdateRequest struct {
    model.Params

    // 商品id
    itemId   int64 

    // 商品数据
    schemaXmlFields   string 

}

func NewAlitripItemSchemaUpdateRequest() *AlitripItemSchemaUpdateRequest{
    return &AlitripItemSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripItemSchemaUpdateRequest) GetApiMethodName() string {
    return "alitrip.item.schema.update"
}

func (r AlitripItemSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripItemSchemaUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlitripItemSchemaUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlitripItemSchemaUpdateRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

func (r AlitripItemSchemaUpdateRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}

