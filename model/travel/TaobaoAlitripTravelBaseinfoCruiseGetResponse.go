package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 APIResponse
taobao.alitrip.travel.baseinfo.cruise.get

旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。
*/
type TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTravelBaseinfoCruiseGetResponse `json:"alitrip_travel_baseinfo_cruise_get_response,omitempty"` 
    TaobaoAlitripTravelBaseinfoCruiseGetResponse
}

/* model for simplify = false
type TaobaoAlitripTravelBaseinfoCruiseGetResponse struct {

    // 邮轮类目扩展信息的json格式字符串
    
    CruiseExtInfos   string `json:"cruise_ext_infos,omitempty"`
    

}
*/

type TaobaoAlitripTravelBaseinfoCruiseGetResponse struct {

    // 邮轮类目扩展信息的json格式字符串
    CruiseExtInfos   string `json:"cruise_ext_infos,omitempty"`

}
