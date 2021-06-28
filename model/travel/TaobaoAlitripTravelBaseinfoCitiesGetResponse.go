package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】度假线路商品发布时基础信息获取接口：地址数据查询 APIResponse
taobao.alitrip.travel.baseinfo.cities.get

旅行度假新商品发布时可用的扩展接口，用于获取可用的出发地或目的地城市列表。
*/
type TaobaoAlitripTravelBaseinfoCitiesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTravelBaseinfoCitiesGetResponse `json:"alitrip_travel_baseinfo_cities_get_response,omitempty"` 
    TaobaoAlitripTravelBaseinfoCitiesGetResponse
}

/* model for simplify = false
type TaobaoAlitripTravelBaseinfoCitiesGetResponse struct {

    // 地区级联城市列表，返回数据为json数组结构的字符串
    
    IocInfos   string `json:"ioc_infos,omitempty"`
    

}
*/

type TaobaoAlitripTravelBaseinfoCitiesGetResponse struct {

    // 地区级联城市列表，返回数据为json数组结构的字符串
    IocInfos   string `json:"ioc_infos,omitempty"`

}
