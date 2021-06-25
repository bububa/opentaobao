package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品编辑接口 APIRequest
tmall.item.vip.schema.update

大商家编辑商品
*/
type TmallItemVipSchemaUpdateRequest struct {
    model.Params

    // 商品编辑的schema参数
    schemaXmlFields   string 

    // 商品id
    itemId   int64 

}

func NewTmallItemVipSchemaUpdateRequest() *TmallItemVipSchemaUpdateRequest{
    return &TmallItemVipSchemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemVipSchemaUpdateRequest) GetApiMethodName() string {
    return "tmall.item.vip.schema.update"
}

func (r TmallItemVipSchemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemVipSchemaUpdateRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

func (r TmallItemVipSchemaUpdateRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}

func (r *TmallItemVipSchemaUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemVipSchemaUpdateRequest) GetItemId() int64 {
    return r.itemId
}

