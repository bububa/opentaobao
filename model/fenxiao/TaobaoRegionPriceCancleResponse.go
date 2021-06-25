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
    Response *TaobaoRegionPriceCancleResponse `json:"taobao_region_price_cancle_response,omitempty"`
}

type TaobaoRegionPriceCancleResponse struct {

    // result
    Result   *BaseResult `json:"result,omitempty"`

}
