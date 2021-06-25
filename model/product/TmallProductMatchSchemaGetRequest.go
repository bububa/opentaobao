package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取匹配产品规则 APIRequest
tmall.product.match.schema.get

ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
*/
type TmallProductMatchSchemaGetRequest struct {
    model.Params

    // 商品发布的目标类目，必须是叶子类目
    categoryId   int64 

}

func NewTmallProductMatchSchemaGetRequest() *TmallProductMatchSchemaGetRequest{
    return &TmallProductMatchSchemaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallProductMatchSchemaGetRequest) GetApiMethodName() string {
    return "tmall.product.match.schema.get"
}

func (r TmallProductMatchSchemaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallProductMatchSchemaGetRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r TmallProductMatchSchemaGetRequest) GetCategoryId() int64 {
    return r.categoryId
}

