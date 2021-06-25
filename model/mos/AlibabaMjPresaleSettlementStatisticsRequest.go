package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预购结算数据统计 APIRequest
alibaba.mj.presale.settlement.statistics

预购结算数据统计
*/
type AlibabaMjPresaleSettlementStatisticsRequest struct {
    model.Params

    // 活动期号
    actionNo   int64 

    // 外部门店编码
    storeNo   string 

}

func NewAlibabaMjPresaleSettlementStatisticsRequest() *AlibabaMjPresaleSettlementStatisticsRequest{
    return &AlibabaMjPresaleSettlementStatisticsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjPresaleSettlementStatisticsRequest) GetApiMethodName() string {
    return "alibaba.mj.presale.settlement.statistics"
}

func (r AlibabaMjPresaleSettlementStatisticsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjPresaleSettlementStatisticsRequest) SetActionNo(actionNo int64) error {
    r.actionNo = actionNo
    r.Set("action_no", actionNo)
    return nil
}

func (r AlibabaMjPresaleSettlementStatisticsRequest) GetActionNo() int64 {
    return r.actionNo
}

func (r *AlibabaMjPresaleSettlementStatisticsRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

func (r AlibabaMjPresaleSettlementStatisticsRequest) GetStoreNo() string {
    return r.storeNo
}

