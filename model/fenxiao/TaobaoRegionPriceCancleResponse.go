package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消区域价格 APIResponse
taobao.region.price.cancle

取消区域价格
*/
type TaobaoRegionPriceCancleAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRegionPriceCancleResponse `json:"region_price_cancle_response,omitempty"` 
    TaobaoRegionPriceCancleResponse
}

/* model for simplify = false
type TaobaoRegionPriceCancleResponse struct {

    // result
    
    Result  *struct {
        BaseResult  *BaseResult `json:"base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoRegionPriceCancleResponse struct {

    // result
    Result   *BaseResult `json:"result,omitempty"`

}
