package ju

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚划算商品搜索接口 APIRequest
taobao.ju.items.search

搜索聚划算商品
*/
type TaobaoJuItemsSearchRequest struct {
    model.Params

    // query
    paramTopItemQuery   *TopItemQuery 

}

func NewTaobaoJuItemsSearchRequest() *TaobaoJuItemsSearchRequest{
    return &TaobaoJuItemsSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJuItemsSearchRequest) GetApiMethodName() string {
    return "taobao.ju.items.search"
}

func (r TaobaoJuItemsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJuItemsSearchRequest) SetParamTopItemQuery(paramTopItemQuery *TopItemQuery) error {
    r.paramTopItemQuery = paramTopItemQuery
    r.Set("param_top_item_query", paramTopItemQuery)
    return nil
}

func (r TaobaoJuItemsSearchRequest) GetParamTopItemQuery() *TopItemQuery {
    return r.paramTopItemQuery
}

