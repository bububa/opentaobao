package crm

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建积分兑换活动 APIResponse
taobao.crm.exchange.activity.create

创建针对积分兑换类型的活动
*/
type TaobaoCrmExchangeActivityCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoCrmExchangeActivityCreateResponse `json:"crm_exchange_activity_create_response,omitempty"` 
    TaobaoCrmExchangeActivityCreateResponse
}

/* model for simplify = false
type TaobaoCrmExchangeActivityCreateResponse struct {

    // 活动ID
    
    ActivityId   int64 `json:"activity_id,omitempty"`
    

    // 人群实例ID
    
    CrowdinstanceId   int64 `json:"crowdinstance_id,omitempty"`
    

    // 接口调用成功
    
    SubSuccess   bool `json:"sub_success,omitempty"`
    

}
*/

type TaobaoCrmExchangeActivityCreateResponse struct {

    // 活动ID
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 人群实例ID
    CrowdinstanceId   int64 `json:"crowdinstance_id,omitempty"`

    // 接口调用成功
    SubSuccess   bool `json:"sub_success,omitempty"`

}
