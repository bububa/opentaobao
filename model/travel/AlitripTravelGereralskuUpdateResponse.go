package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发布SKU信息（如果properties重复 则更新） APIResponse
alitrip.travel.gereralsku.update

发布SKU信息（如果properties重复 则更新）
*/
type AlitripTravelGereralskuUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlitripTravelGereralskuUpdateResponse `json:"alitrip_travel_gereralsku_update_response,omitempty"`
}

type AlitripTravelGereralskuUpdateResponse struct {

    // 返回结果
    FirstResult   *TopTravelItem `json:"first_result,omitempty"`

}
