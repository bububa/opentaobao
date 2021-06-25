package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫简化发布商品 APIRequest
tmall.item.simpleschema.add

天猫简化版schema发布商品。
*/
type TmallItemSimpleschemaAddRequest struct {
    model.Params

    // 根据tmall.item.add.simpleschema.get生成的商品发布规则入参数据
    schemaXmlFields   string 

}

func NewTmallItemSimpleschemaAddRequest() *TmallItemSimpleschemaAddRequest{
    return &TmallItemSimpleschemaAddRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemSimpleschemaAddRequest) GetApiMethodName() string {
    return "tmall.item.simpleschema.add"
}

func (r TmallItemSimpleschemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemSimpleschemaAddRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

func (r TmallItemSimpleschemaAddRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}

