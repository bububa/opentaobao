package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品编辑接口 APIResponse
taobao.alitrip.travel.item.base.modify

旅行度假新商品基本信息修改接口 第三版。提供商家通过TOP API方式修改商品除sku外的基本信息。
*/
type TaobaoAlitripTravelItemBaseModifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelItemBaseModifyResponse `json:"taobao_alitrip_travel_item_base_modify_response,omitempty"`
}

type TaobaoAlitripTravelItemBaseModifyResponse struct {

    // 商品修改结果
    TravelItem   *TopTravelItem `json:"travel_item,omitempty"`

}
