package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaJymRequirementOrdersOperationNotifyAPIRequest
交易猫需求订单操作接口 API请求
alibaba.jym.requirement.orders.operation.notify

交易猫需求订单操作接口 */
type AlibabaJymRequirementOrdersOperationNotifyAPIRequest struct {
	model.Params
	// 需求订单操作
	_operation int64
	// 需求订单id
	_reqmntOrderId string
}

// New
