package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
区域价格查询 APIResponse
taobao.region.price.query

区域价格查询
*/
type TaobaoRegionPriceQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRegionPriceQueryResponse `json:"region_price_query_response,omitempty"` 
    TaobaoRegionPriceQueryResponse
}

/* model for simplify = false
type TaobaoRegionPriceQueryResponse struct {

    // result
    
    Result  *struct {
        BaseResult  *BaseResult `json:"base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRegionPriceQueryResponse struct {

    // result
    Result   *BaseResult `json:"result,omitempty"`

}
