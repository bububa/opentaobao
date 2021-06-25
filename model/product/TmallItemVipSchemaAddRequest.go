package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大商家商品发布接口 APIRequest
tmall.item.vip.schema.add

大商家商品发布接口
*/
type TmallItemVipSchemaAddRequest struct {
    model.Params

    // 商品发布schema参数
    schemaXmlFields   string 

}

func NewTmallItemVipSchemaAddRequest() *TmallItemVipSchemaAddRequest{
    return &TmallItemVipSchemaAddRequest{
        Params: model.NewParams(),
    }
}

func (r TmallItemVipSchemaAddRequest) GetApiMethodName() string {
    return "tmall.item.vip.schema.add"
}

func (r TmallItemVipSchemaAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallItemVipSchemaAddRequest) SetSchemaXmlFields(schemaXmlFields string) error {
    r.schemaXmlFields = schemaXmlFields
    r.Set("schema_xml_fields", schemaXmlFields)
    return nil
}

func (r TmallItemVipSchemaAddRequest) GetSchemaXmlFields() string {
    return r.schemaXmlFields
}

