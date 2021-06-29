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
type TmallServicecenterWorkcardExpressorderCreateRequest struct {
    model.Params
    // 物流单关联的工单List
    workcardIdList   []int64
    // 真实履约服务商Nick（非ERP系统不要填写）
    realTpNick   string
    // erp外部物流订单号
    externalLogisticsOrderId   string
}

// 初始化TmallServicecenterWorkcardExpressorderCreateRequest对象
func NewTmallServicecenterWorkcardExpressorderCreateRequest() *TmallServicecenterWorkcardExpressorderCreateRequest{
    return &TmallServicecenterWorkcardExpressorderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardExpressorderCreateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.expressorder.create"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardExpressorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardIdList Setter
// 物流单关联的工单List
func (r *TmallServicecenterWorkcardExpressorderCreateRequest) SetWorkcardIdList(workcardIdList []int64) error {
    r.workcardIdList = workcardIdList
    r.Set("workcard_id_list", workcardIdList)
    return nil
}

// WorkcardIdList Getter
func (r TmallServicecenterWorkcardExpressorderCreateRequest) GetWorkcardIdList() []int64 {
    return r.workcardIdList
}
// RealTpNick Setter
// 真实履约服务商Nick（非ERP系统不要填写）
func (r *TmallServicecenterWorkcardExpressorderCreateRequest) SetRealTpNick(realTpNick string) error {
    r.realTpNick = realTpNick
    r.Set("real_tp_nick", realTpNick)
    return nil
}

// RealTpNick Getter
func (r TmallServicecenterWorkcardExpressorderCreateRequest) GetRealTpNick() string {
    return r.realTpNick
}
// ExternalLogisticsOrderId Setter
// erp外部物流订单号
func (r *TmallServicecenterWorkcardExpressorderCreateRequest) SetExternalLogisticsOrderId(externalLogisticsOrderId string) error {
    r.externalLogisticsOrderId = externalLogisticsOrderId
    r.Set("external_logistics_order_id", externalLogisticsOrderId)
    return nil
}

// ExternalLogisticsOrderId Getter
func (r TmallServicecenterWorkcardExpressorderCreateRequest) GetExternalLogisticsOrderId() string {
    return r.externalLogisticsOrderId
}
