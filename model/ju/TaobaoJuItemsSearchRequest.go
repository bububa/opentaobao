package ju

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚划算商品搜索接口 API请求
taobao.ju.items.search

搜索聚划算商品
*/
type TaobaoJuItemsSearchAPIRequest struct {
    model.Params
    // query
    _paramTopItemQuery   *TopItemQuery
}

// 初始化TaobaoJuItemsSearchAPIRequest对象
func NewTaobaoJuItemsSearchRequest() *TaobaoJuItemsSearchAPIRequest{
    return &TaobaoJuItemsSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJuItemsSearchAPIRequest) GetApiMethodName() string {
    return "taobao.ju.items.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJuItemsSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopItemQuery Setter
// query
func (r *TaobaoJuItemsSearchAPIRequest) SetParamTopItemQuery(_paramTopItemQuery *TopItemQuery) error {
    r._paramTopItemQuery = _paramTopItemQuery
    r.Set("param_top_item_query", _paramTopItemQuery)
    return nil
}

// ParamTopItemQuery Getter
func (r TaobaoJuItemsSearchAPIRequest) GetParamTopItemQuery() *TopItemQuery {
    return r._paramTopItemQuery
}
