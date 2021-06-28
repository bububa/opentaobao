package qt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝服务订购关系查询 APIResponse
taobao.ts.subscribe.get

ts订购关系状态查询. 暂只支持1口价服务.
*/
type TaobaoTsSubscribeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTsSubscribeGetResponse `json:"ts_subscribe_get_response,omitempty"` 
    TaobaoTsSubscribeGetResponse
}

/* model for simplify = false
type TaobaoTsSubscribeGetResponse struct {

    // 订购关系对象
    
    ServiceSubscribe  *struct {
        ServiceSubscribe  *ServiceSubscribe `json:"service_subscribe,omitempty"`
    } `json:"service_subscribe,omitempty"`
    

}
*/

type TaobaoTsSubscribeGetResponse struct {

    // 订购关系对象
    ServiceSubscribe   *ServiceSubscribe `json:"service_subscribe,omitempty"`

}
