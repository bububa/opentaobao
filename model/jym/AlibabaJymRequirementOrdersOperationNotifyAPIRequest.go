package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymRequirementOrdersOperationNotifyAPIRequest 交易猫需求订单操作接口 API请求
// alibaba.jym.requirement.orders.operation.notify
//
// 交易猫需求订单操作接口
type AlibabaJymRequirementOrdersOperationNotifyAPIRequest struct {
	model.Params
	// 需求订单id
	_reqmntOrderId string
	// 需求订单操作
	_operation int64
}

// NewAlibabaJymRequirementOrdersOperationNotifyRequest 初始化AlibabaJymRequirementOrdersOperationNotifyAPIRequest对象
func NewAlibabaJymRequirementOrdersOperationNotifyRequest() *AlibabaJymRequirementOrdersOperationNotifyAPIRequest {
	return &AlibabaJymRequirementOrdersOperationNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymRequirementOrdersOperationNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.requirement.orders.operation.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymRequirementOrdersOperationNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReqmntOrderId is ReqmntOrderId Setter
// 需求订单id
func (r *AlibabaJymRequirementOrdersOperationNotifyAPIRequest) SetReqmntOrderId(_reqmntOrderId string) error {
	r._reqmntOrderId = _reqmntOrderId
	r.Set("reqmnt_order_id", _reqmntOrderId)
	return nil
}

// GetReqmntOrderId ReqmntOrderId Getter
func (r AlibabaJymRequirementOrdersOperationNotifyAPIRequest) GetReqmntOrderId() string {
	return r._reqmntOrderId
}

// SetOperation is Operation Setter
// 需求订单操作
func (r *AlibabaJymRequirementOrdersOperationNotifyAPIRequest) SetOperation(_operation int64) error {
	r._operation = _operation
	r.Set("operation", _operation)
	return nil
}

// GetOperation Operation Getter
func (r AlibabaJymRequirementOrdersOperationNotifyAPIRequest) GetOperation() int64 {
	return r._operation
}
