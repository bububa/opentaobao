package crm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建积分兑换活动 API返回值 
taobao.crm.exchange.activity.create

创建针对积分兑换类型的活动
*/
type TaobaoCrmExchangeActivityCreateAPIResponse struct {
    model.CommonResponse
    TaobaoCrmExchangeActivityCreateAPIResponseModel
}

// 创建积分兑换活动 成功返回结果
type TaobaoCrmExchangeActivityCreateAPIResponseModel struct {
    XMLName xml.Name `xml:"crm_exchange_activity_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 活动ID
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 人群实例ID
    CrowdinstanceId   int64 `json:"crowdinstance_id,omitempty" xml:"crowdinstance_id,omitempty"`
    // 接口调用成功
    SubSuccess   bool `json:"sub_success,omitempty" xml:"sub_success,omitempty"`
}
