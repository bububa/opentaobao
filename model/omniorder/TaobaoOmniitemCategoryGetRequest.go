package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道商品轻发布类目信息 APIRequest
taobao.omniitem.category.get

全渠道商品轻发布类目信息
*/
type TaobaoOmniitemCategoryGetRequest struct {
    model.Params

    // 全渠道商品类目ID，不填表示获取所有全渠道商品类目信息
    categoryId   int64 

}

func NewTaobaoOmniitemCategoryGetRequest() *TaobaoOmniitemCategoryGetRequest{
    return &TaobaoOmniitemCategoryGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniitemCategoryGetRequest) GetApiMethodName() string {
    return "taobao.omniitem.category.get"
}

func (r TaobaoOmniitemCategoryGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniitemCategoryGetRequest) SetCategoryId(categoryId int64) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r TaobaoOmniitemCategoryGetRequest) GetCategoryId() int64 {
    return r.categoryId
}

