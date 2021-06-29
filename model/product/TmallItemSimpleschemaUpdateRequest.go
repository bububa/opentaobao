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
    _itemId   int64
    // 编辑商品时提交的xml信息
    _schemaXmlFields   string
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
func (r *TmallItemSimpleschemaUpdateRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TmallItemSimpleschemaUpdateRequest) GetItemId() int64 {
    return r._itemId
}
// SchemaXmlFields Setter
// 编辑商品时提交的xml信息
func (r *TmallItemSimpleschemaUpdateRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
    r._schemaXmlFields = _schemaXmlFields
    r.Set("schema_xml_fields", _schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r TmallItemSimpleschemaUpdateRequest) GetSchemaXmlFields() string {
    return r._schemaXmlFields
}
