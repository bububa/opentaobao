package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品发布接口 API请求
tmall.item.vip.schema.add

大商家商品发布接口
*/
type TmallItemVipSchemaAddRequest struct {
    model.Params
    // 商品发布schema参数
    _schemaXmlFields   string
}

// 初始化TmallItemVipSchemaAddRequest对象
func NewTmallItemVipSchemaAddRequest() *TmallItemVipSchemaAddRequest{
    return &TmallItemVipSchemaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemVipSchemaAddRequest) GetApiMethodName() string {
    return "tmall.item.vip.schema.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemVipSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SchemaXmlFields Setter
// 商品发布schema参数
func (r *TmallItemVipSchemaAddRequest) SetSchemaXmlFields(_schemaXmlFields string) error {
    r._schemaXmlFields = _schemaXmlFields
    r.Set("schema_xml_fields", _schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r TmallItemVipSchemaAddRequest) GetSchemaXmlFields() string {
    return r._schemaXmlFields
}
