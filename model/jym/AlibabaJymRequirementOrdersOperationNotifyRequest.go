package jym

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易猫需求订单操作接口 API请求
alibaba.jym.requirement.orders.operation.notify

交易猫需求订单操作接口
*/
type AlibabaJymRequirementOrdersOperationNotifyRequest struct {
    model.Params
    // 需求订单操作
    _operation   int64
    // 需求订单id
    _reqmntOrderId   string
}

// 初始化AlibabaJymRequirementOrdersOperationNotifyRequest对象
func NewAlibabaJymRequirementOrdersOperationNotifyRequest() *AlibabaJymRequirementOrdersOperationNotifyRequest{
    return &AlibabaJymRequirementOrdersOperationNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetApiMethodName() string {
    return "alibaba.jym.requirement.orders.operation.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Operation Setter
// 需求订单操作
func (r *AlibabaJymRequirementOrdersOperationNotifyRequest) SetOperation(_operation int64) error {
    r._operation = _operation
    r.Set("operation", _operation)
    return nil
}

// Operation Getter
func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetOperation() int64 {
    return r._operation
}
// ReqmntOrderId Setter
// 需求订单id
func (r *AlibabaJymRequirementOrdersOperationNotifyRequest) SetReqmntOrderId(_reqmntOrderId string) error {
    r._reqmntOrderId = _reqmntOrderId
    r.Set("reqmnt_order_id", _reqmntOrderId)
    return nil
}

// ReqmntOrderId Getter
func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetReqmntOrderId() string {
    return r._reqmntOrderId
}
