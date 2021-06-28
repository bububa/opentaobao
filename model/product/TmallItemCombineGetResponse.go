package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合商品获取接口 APIResponse
tmall.item.combine.get

查询组合商品的SKU信息
*/
type TmallItemCombineGetAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemCombineGetResponse `json:"tmall_item_combine_get_response,omitempty"` 
    TmallItemCombineGetResponse
}

/* model for simplify = false
type TmallItemCombineGetResponse struct {

    // results
    
    Results  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TmallItemCombineGetResponse struct {

    // results
    Results   []string `json:"results,omitempty"`

}
