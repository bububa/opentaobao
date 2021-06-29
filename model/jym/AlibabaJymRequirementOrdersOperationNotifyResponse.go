package jym

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
交易猫需求订单操作接口 APIResponse
alibaba.jym.requirement.orders.operation.notify

交易猫需求订单操作接口
*/
type AlibabaJymRequirementOrdersOperationNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaJymRequirementOrdersOperationNotifyResponse
}

type AlibabaJymRequirementOrdersOperationNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_jym_requirement_orders_operation_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    ResultDto   *AlibabaJymRequirementOrdersOperationNotifyResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`

    
}
