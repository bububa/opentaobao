package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
(供销)船票通用类目sku新增&编辑API APIResponse
alitrip.travel.product.gereralsku.update

发布SKU信息（如果properties重复 则更新）
*/
type AlitripTravelProductGereralskuUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlitripTravelProductGereralskuUpdateResponse `json:"alitrip_travel_product_gereralsku_update_response,omitempty"`
}

type AlitripTravelProductGereralskuUpdateResponse struct {

    // 返回结果
    FirstResult   *TopTravelItem `json:"first_result,omitempty"`

}
