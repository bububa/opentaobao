package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫开新车租后异常流线下处理状态通知接口 API请求
tmall.car.lease.exceptionflowsynchronize

天猫开新车租后异常流线下处理状态通知接口
*/
type TmallCarLeaseExceptionflowsynchronizeRequest struct {
    model.Params
    // 天猫开新车订单id
    orderId   int64
    // 1:开始切换为异常流，2:线下处理完成
    status   int64
    // 异常流类型,0.退车,1.买断,2.分期，3.续租
    flowType   int64
    // 切换原因描述
    desc   string
}

// 初始化TmallCarLeaseExceptionflowsynchronizeRequest对象
func NewTmallCarLeaseExceptionflowsynchronizeRequest() *TmallCarLeaseExceptionflowsynchronizeRequest{
    return &TmallCarLeaseExceptionflowsynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseExceptionflowsynchronizeRequest) GetApiMethodName() string {
    return "tmall.car.lease.exceptionflowsynchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseExceptionflowsynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 天猫开新车订单id
func (r *TmallCarLeaseExceptionflowsynchronizeRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeaseExceptionflowsynchronizeRequest) GetOrderId() int64 {
    return r.orderId
}
// Status Setter
// 1:开始切换为异常流，2:线下处理完成
func (r *TmallCarLeaseExceptionflowsynchronizeRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TmallCarLeaseExceptionflowsynchronizeRequest) GetStatus() int64 {
    return r.status
}
// FlowType Setter
// 异常流类型,0.退车,1.买断,2.分期，3.续租
func (r *TmallCarLeaseExceptionflowsynchronizeRequest) SetFlowType(flowType int64) error {
    r.flowType = flowType
    r.Set("flow_type", flowType)
    return nil
}

// FlowType Getter
func (r TmallCarLeaseExceptionflowsynchronizeRequest) GetFlowType() int64 {
    return r.flowType
}
// Desc Setter
// 切换原因描述
func (r *TmallCarLeaseExceptionflowsynchronizeRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TmallCarLeaseExceptionflowsynchronizeRequest) GetDesc() string {
    return r.desc
}
