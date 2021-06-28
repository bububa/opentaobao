package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
判定服务是否可达 APIResponse
taobao.logistics.address.reachable

根据输入的目标地址，判断服务是否可达。
现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
*/
type TaobaoLogisticsAddressReachableAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsAddressReachableResponse `json:"logistics_address_reachable_response,omitempty"` 
    TaobaoLogisticsAddressReachableResponse
}

/* model for simplify = false
type TaobaoLogisticsAddressReachableResponse struct {

    // 地址可达返回结果，每个TP对应一个
    
    ReachableResultList  struct {
        AddressReachableResult  []AddressReachableResult `json:"address_reachable_result,omitempty"`
    } `json:"reachable_result_list,omitempty"`
    

}
*/

type TaobaoLogisticsAddressReachableResponse struct {

    // 地址可达返回结果，每个TP对应一个
    ReachableResultList   []AddressReachableResult `json:"reachable_result_list,omitempty"`

}
