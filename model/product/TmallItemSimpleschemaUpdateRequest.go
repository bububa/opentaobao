package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化编辑商品 API请求
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

// 初始化TmallItemSimpleschemaUpdateRequest对象
func NewTmallItemSimpleschemaUpdateRequest() *TmallItemSimpleschemaUpdateRequest{
    return &TmallItemSimpleschemaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSimpleschemaUpdateRequest) GetApiMethodName() string {
    return "tmall.item.simpleschema.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSimpleschemaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemId Setter
// 商品id
func (r *TmallItemSimpleschemaUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TmallItemSimpleschemaUpdateRequest) GetItemId() int64 {
    return r.itemId
}
// SchemaXmlFields Setter
// 编辑商品时提交的xml信息
func (r *TmallItemSimpleschemaUpdateRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r TmallItemSimpleschemaUpdateRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}
