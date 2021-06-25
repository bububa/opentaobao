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
    Response *TaobaoRegionPriceManageResponse `json:"taobao_region_price_manage_response,omitempty"`
}

type TaobaoRegionPriceManageResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
