package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化编辑商品 APIRequest
tmall.item.simpleschema.update

国外大商家天猫简化编辑商品
*/
type TmallItemSimpleschemaUpdateRequest struct {
    model.Params

    // 商品id
    itemId   int64 

    // 编辑商品时提交的xml信息
    schemaXmlFields   string 

}

func NewTmallItemSimpleschemaUpdateRequest() *TmallItemSimpleschemaUpdateRequest{
    return &TmallItemSimpleschemaUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSimpleschemaUpdateRequest) GetApiMethodName() string {
    return "tmall.item.simpleschema.update"
}

func (r TmallItemSimpleschemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSimpleschemaUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TmallItemSimpleschemaUpdateRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TmallItemSimpleschemaUpdateRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

func (r TmallItemSimpleschemaUpdateRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}

