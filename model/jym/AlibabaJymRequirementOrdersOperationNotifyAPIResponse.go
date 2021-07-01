package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaJymRequirementOrdersOperationNotifyAPIResponse
交易猫需求订单操作接口 API返回值
alibaba.jym.requirement.orders.operation.notify

交易猫需求订单操作接口 */
type AlibabaJymRequirementOrdersOperationNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaJymRequirementOrdersOperationNotifyAPIResponseModel
}

// AlibabaJymRequirementOrdersOperationNotifyAPIResponseModel is 交易猫需求订单操作接口 成功返回结果
type AlibabaJymRequirementOrdersOperationNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_requirement_orders_operation_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	ResultDto *AlibabaJymRequirementOrdersOperationNotifyResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}
