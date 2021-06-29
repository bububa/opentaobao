package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单呼叫运力 API请求
tmall.servicecenter.workcard.expressorder.consign

天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
*/
type TmallServicecenterWorkcardExpressorderConsignRequest struct {
    model.Params
    // 物流寄件单号（废弃）
    expressOrderId   int64
    // 工单List
    workcardIdList   []int64
    // 真实接单服务商
    realTpNick   string
    // 物流订单号
    logisticsOrderId   int64
}

// 初始化TmallServicecenterWorkcardExpressorderConsignRequest对象
func NewTmallServicecenterWorkcardExpressorderConsignRequest() *TmallServicecenterWorkcardExpressorderConsignRequest{
    return &TmallServicecenterWorkcardExpressorderConsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.expressorder.consign"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExpressOrderId Setter
// 物流寄件单号（废弃）
func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetExpressOrderId(expressOrderId int64) error {
    r.expressOrderId = expressOrderId
    r.Set("express_order_id", expressOrderId)
    return nil
}

// ExpressOrderId Getter
func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetExpressOrderId() int64 {
    return r.expressOrderId
}
// WorkcardIdList Setter
// 工单List
func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetWorkcardIdList(workcardIdList []int64) error {
    r.workcardIdList = workcardIdList
    r.Set("workcard_id_list", workcardIdList)
    return nil
}

// WorkcardIdList Getter
func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetWorkcardIdList() []int64 {
    return r.workcardIdList
}
// RealTpNick Setter
// 真实接单服务商
func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetRealTpNick(realTpNick string) error {
    r.realTpNick = realTpNick
    r.Set("real_tp_nick", realTpNick)
    return nil
}

// RealTpNick Getter
func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetRealTpNick() string {
    return r.realTpNick
}
// LogisticsOrderId Setter
// 物流订单号
func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetLogisticsOrderId(logisticsOrderId int64) error {
    r.logisticsOrderId = logisticsOrderId
    r.Set("logistics_order_id", logisticsOrderId)
    return nil
}

// LogisticsOrderId Getter
func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetLogisticsOrderId() int64 {
    return r.logisticsOrderId
}
