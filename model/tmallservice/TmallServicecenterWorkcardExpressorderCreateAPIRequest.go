package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单创建API API请求
tmall.servicecenter.workcard.expressorder.create

天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API
*/
type TmallServicecenterWorkcardExpressorderCreateAPIRequest struct {
    model.Params
    // 物流单关联的工单List
    _workcardIdList   []int64
    // 真实履约服务商Nick（非ERP系统不要填写）
    _realTpNick   string
    // erp外部物流订单号
    _externalLogisticsOrderId   string
}

// 初始化TmallServicecenterWorkcardExpressorderCreateAPIRequest对象
func NewTmallServicecenterWorkcardExpressorderCreateRequest() *TmallServicecenterWorkcardExpressorderCreateAPIRequest{
    return &TmallServicecenterWorkcardExpressorderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardExpressorderCreateAPIRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.expressorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardExpressorderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardIdList Setter
// 物流单关联的工单List
func (r *TmallServicecenterWorkcardExpressorderCreateAPIRequest) SetWorkcardIdList(_workcardIdList []int64) error {
    r._workcardIdList = _workcardIdList
    r.Set("workcard_id_list", _workcardIdList)
    return nil
}

// WorkcardIdList Getter
func (r TmallServicecenterWorkcardExpressorderCreateAPIRequest) GetWorkcardIdList() []int64 {
    return r._workcardIdList
}
// RealTpNick Setter
// 真实履约服务商Nick（非ERP系统不要填写）
func (r *TmallServicecenterWorkcardExpressorderCreateAPIRequest) SetRealTpNick(_realTpNick string) error {
    r._realTpNick = _realTpNick
    r.Set("real_tp_nick", _realTpNick)
    return nil
}

// RealTpNick Getter
func (r TmallServicecenterWorkcardExpressorderCreateAPIRequest) GetRealTpNick() string {
    return r._realTpNick
}
// ExternalLogisticsOrderId Setter
// erp外部物流订单号
func (r *TmallServicecenterWorkcardExpressorderCreateAPIRequest) SetExternalLogisticsOrderId(_externalLogisticsOrderId string) error {
    r._externalLogisticsOrderId = _externalLogisticsOrderId
    r.Set("external_logistics_order_id", _externalLogisticsOrderId)
    return nil
}

// ExternalLogisticsOrderId Getter
func (r TmallServicecenterWorkcardExpressorderCreateAPIRequest) GetExternalLogisticsOrderId() string {
    return r._externalLogisticsOrderId
}
