package ju

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ju"
)

/* 
聚划算商品搜索接口 
taobao.ju.items.search

搜索聚划算商品
*/
func TaobaoJuItemsSearch(clt *core.SDKClient, req *ju.TaobaoJuItemsSearchRequest, session string) (*ju.TaobaoJuItemsSearchAPIResponse, error) {
    var resp ju.TaobaoJuItemsSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
