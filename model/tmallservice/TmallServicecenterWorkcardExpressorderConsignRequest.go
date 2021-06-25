package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流订单呼叫运力 APIRequest
tmall.servicecenter.workcard.expressorder.consign

天猫服务寄送类业务，服务商履约完成后进行寄回操作呼叫运力
*/
type TmallServicecenterWorkcardExpressorderConsignRequest struct {
    model.Params

    // 物流寄件单号（废弃）
    expressOrderId   int64 

    // 工单List
    workcardIdList   []Number 

    // 真实接单服务商
    realTpNick   string 

    // 物流订单号
    logisticsOrderId   int64 

}

func NewTmallServicecenterWorkcardExpressorderConsignRequest() *TmallServicecenterWorkcardExpressorderConsignRequest{
    return &TmallServicecenterWorkcardExpressorderConsignRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.expressorder.consign"
}

func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetExpressOrderId(expressOrderId int64) error {
    r.expressOrderId = expressOrderId
    r.Set("express_order_id", expressOrderId)
    return nil
}

func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetExpressOrderId() int64 {
    return r.expressOrderId
}

func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetWorkcardIdList(workcardIdList []Number) error {
    r.workcardIdList = workcardIdList
    r.Set("workcard_id_list", workcardIdList)
    return nil
}

func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetWorkcardIdList() []Number {
    return r.workcardIdList
}

func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetRealTpNick(realTpNick string) error {
    r.realTpNick = realTpNick
    r.Set("real_tp_nick", realTpNick)
    return nil
}

func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetRealTpNick() string {
    return r.realTpNick
}

func (r *TmallServicecenterWorkcardExpressorderConsignRequest) SetLogisticsOrderId(logisticsOrderId int64) error {
    r.logisticsOrderId = logisticsOrderId
    r.Set("logistics_order_id", logisticsOrderId)
    return nil
}

func (r TmallServicecenterWorkcardExpressorderConsignRequest) GetLogisticsOrderId() int64 {
    return r.logisticsOrderId
}

