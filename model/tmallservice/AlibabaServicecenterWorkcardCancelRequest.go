package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台工单取消接口 API请求
alibaba.servicecenter.workcard.cancel

取消服务工单
*/
type AlibabaServicecenterWorkcardCancelRequest struct {
    model.Params
    // 工单id
    workcardId   int64
    // 取消备注
    memo   string
    // 服务单id
    serviceOrderId   int64
    // 真实服务商昵称
    realTpNick   string
}

// 初始化AlibabaServicecenterWorkcardCancelRequest对象
func NewAlibabaServicecenterWorkcardCancelRequest() *AlibabaServicecenterWorkcardCancelRequest{
    return &AlibabaServicecenterWorkcardCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterWorkcardCancelRequest) GetApiMethodName() string {
    return "alibaba.servicecenter.workcard.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterWorkcardCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardId Setter
// 工单id
func (r *AlibabaServicecenterWorkcardCancelRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetWorkcardId() int64 {
    return r.workcardId
}
// Memo Setter
// 取消备注
func (r *AlibabaServicecenterWorkcardCancelRequest) SetMemo(memo string) error {
    r.memo = memo
    r.Set("memo", memo)
    return nil
}

// Memo Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetMemo() string {
    return r.memo
}
// ServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterWorkcardCancelRequest) SetServiceOrderId(serviceOrderId int64) error {
    r.serviceOrderId = serviceOrderId
    r.Set("service_order_id", serviceOrderId)
    return nil
}

// ServiceOrderId Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetServiceOrderId() int64 {
    return r.serviceOrderId
}
// RealTpNick Setter
// 真实服务商昵称
func (r *AlibabaServicecenterWorkcardCancelRequest) SetRealTpNick(realTpNick string) error {
    r.realTpNick = realTpNick
    r.Set("real_tp_nick", realTpNick)
    return nil
}

// RealTpNick Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetRealTpNick() string {
    return r.realTpNick
}
