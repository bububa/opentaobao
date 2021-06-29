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
    _workcardId   int64
    // 取消备注
    _memo   string
    // 服务单id
    _serviceOrderId   int64
    // 真实服务商昵称
    _realTpNick   string
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
func (r *AlibabaServicecenterWorkcardCancelRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetWorkcardId() int64 {
    return r._workcardId
}
// Memo Setter
// 取消备注
func (r *AlibabaServicecenterWorkcardCancelRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetMemo() string {
    return r._memo
}
// ServiceOrderId Setter
// 服务单id
func (r *AlibabaServicecenterWorkcardCancelRequest) SetServiceOrderId(_serviceOrderId int64) error {
    r._serviceOrderId = _serviceOrderId
    r.Set("service_order_id", _serviceOrderId)
    return nil
}

// ServiceOrderId Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetServiceOrderId() int64 {
    return r._serviceOrderId
}
// RealTpNick Setter
// 真实服务商昵称
func (r *AlibabaServicecenterWorkcardCancelRequest) SetRealTpNick(_realTpNick string) error {
    r._realTpNick = _realTpNick
    r.Set("real_tp_nick", _realTpNick)
    return nil
}

// RealTpNick Getter
func (r AlibabaServicecenterWorkcardCancelRequest) GetRealTpNick() string {
    return r._realTpNick
}
