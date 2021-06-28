package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建积分兑换活动 APIResponse
taobao.crm.exchange.activity.create

创建针对积分兑换类型的活动
*/
type TaobaoCrmExchangeActivityCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"crm_exchange_activity_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 活动ID
    
    ActivityId   int64 `json:"activity_id,omitempty" xml:"