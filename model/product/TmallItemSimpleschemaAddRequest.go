package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化发布商品 API请求
tmall.item.simpleschema.add

天猫简化版schema发布商品。
*/
type TmallItemSimpleschemaAddRequest struct {
    model.Params
    // 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
    schemaXmlFields   string
}

// 初始化TmallItemSimpleschemaAddRequest对象
func NewTmallItemSimpleschemaAddRequest() *TmallItemSimpleschemaAddRequest{
    return &TmallItemSimpleschemaAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallItemSimpleschemaAddRequest) GetApiMethodName() string {
    return "tmall.item.simpleschema.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallItemSimpleschemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SchemaXmlFields Setter
// 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
func (r *TmallItemSimpleschemaAddRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

// SchemaXmlFields Getter
func (r TmallItemSimpleschemaAddRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}
