package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
编辑区域价格 APIResponse
taobao.region.price.manage

编辑区域价格
*/
type TaobaoRegionPriceManageAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRegionPriceManageResponse `json:"region_price_manage_response,omitempty"` 
    TaobaoRegionPriceManageResponse
}

/* model for simplify = false
type TaobaoRegionPriceManageResponse struct {

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoRegionPriceManageResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
