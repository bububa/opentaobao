package ju

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚划算商品搜索接口 APIResponse
taobao.ju.items.search

搜索聚划算商品
*/
type TaobaoJuItemsSearchAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJuItemsSearchResponse `json:"taobao_ju_items_search_response,omitempty"`
}

type TaobaoJuItemsSearchResponse struct {

    // 返回结果
    Result   *PaginationResult `json:"result,omitempty"`

}
