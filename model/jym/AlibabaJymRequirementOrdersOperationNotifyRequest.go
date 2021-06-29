package jym

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
交易猫需求订单操作接口 APIRequest
alibaba.jym.requirement.orders.operation.notify

交易猫需求订单操作接口
*/
type AlibabaJymRequirementOrdersOperationNotifyRequest struct {
    model.Params

    // 需求订单操作
    operation   int64 

    // 需求订单id
    reqmntOrderId   string 

}

func NewAlibabaJymRequirementOrdersOperationNotifyRequest() *AlibabaJymRequirementOrdersOperationNotifyRequest{
    return &AlibabaJymRequirementOrdersOperationNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetApiMethodName() string {
    return "alibaba.jym.requirement.orders.operation.notify"
}

func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaJymRequirementOrdersOperationNotifyRequest) SetOperation(operation int64) error {
    r.operation = operation
    r.Set("operation", operation)
    return nil
}

func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetOperation() int64 {
    return r.operation
}

func (r *AlibabaJymRequirementOrdersOperationNotifyRequest) SetReqmntOrderId(reqmntOrderId string) error {
    r.reqmntOrderId = reqmntOrderId
    r.Set("reqmnt_order_id", reqmntOrderId)
    return nil
}

func (r AlibabaJymRequirementOrdersOperationNotifyRequest) GetReqmntOrderId() string {
    return r.reqmntOrderId
}

