package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-店铺搜索 APIResponse
taobao.tbk.shop.get

淘宝客店铺查询
*/
type TaobaoTbkShopGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTbkShopGetResponse `json:"tbk_shop_get_response,omitempty"` 
    TaobaoTbkShopGetResponse
}

/* model for simplify = false
type TaobaoTbkShopGetResponse struct {

    // 淘宝客店铺
    
    Results  struct {
        NTbkShop  []NTbkShop `json:"n_tbk_shop,omitempty"`
    } `json:"results,omitempty"`
    

    // 搜索到符合条件的结果总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoTbkShopGetResponse struct {

    // 淘宝客店铺
    Results   []NTbkShop `json:"results,omitempty"`

    // 搜索到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
