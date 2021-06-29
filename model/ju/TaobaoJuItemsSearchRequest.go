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
type TaobaoJuItemsSearchRequest struct {
    model.Params
    // query
    _paramTopItemQuery   *TopItemQuery
}

// 初始化TaobaoJuItemsSearchRequest对象
func NewTaobaoJuItemsSearchRequest() *TaobaoJuItemsSearchRequest{
    return &TaobaoJuItemsSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJuItemsSearchRequest) GetApiMethodName() string {
    return "taobao.ju.items.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJuItemsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopItemQuery Setter
// query
func (r *TaobaoJuItemsSearchRequest) SetParamTopItemQuery(_paramTopItemQuery *TopItemQuery) error {
    r._paramTopItemQuery = _paramTopItemQuery
    r.Set("param_top_item_query", _paramTopItemQuery)
    return nil
}

// ParamTopItemQuery Getter
func (r TaobaoJuItemsSearchRequest) GetParamTopItemQuery() *TopItemQuery {
    return r._paramTopItemQuery
}
