package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
product匹配接口 APIRequest
tmall.product.schema.match

根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
*/
type TmallProductSchemaMatchRequest struct {
    model.Params

    // 商品发布的目标类目，必须是叶子类目
    categoryId   int64 

    // 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
    propvalues   string 

}

func NewTmallProductSchemaMatchRequest() *TmallProductSchemaMatchRequest{
    return &TmallProductSchemaMatchRequest{
        Params: model.NewParams(),
    }
}

func (r TmallProductSchemaMatchRequest) GetApiMethodName() string {
    return "tmall.product.schema.match"
}

func (r TmallProductSchemaMatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallProductSchemaMatchRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r TmallProductSchemaMatchRequest) GetCategoryId() int64 {
    return r.categoryId
}

func (r *TmallProductSchemaMatchRequest) SetPropvalues(propvalues string) error {
    r.propvalues = propvalues
    r.Set("propvalues", propvalues)
    return nil
}

func (r TmallProductSchemaMatchRequest) GetPropvalues() string {
    return r.propvalues
}

