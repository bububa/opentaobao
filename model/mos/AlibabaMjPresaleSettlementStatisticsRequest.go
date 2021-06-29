package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
预购结算数据统计 API请求
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

// 初始化AlibabaMjPresaleSettlementStatisticsRequest对象
func NewAlibabaMjPresaleSettlementStatisticsRequest() *AlibabaMjPresaleSettlementStatisticsRequest{
    return &AlibabaMjPresaleSettlementStatisticsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjPresaleSettlementStatisticsRequest) GetApiMethodName() string {
    return "alibaba.mj.presale.settlement.statistics"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjPresaleSettlementStatisticsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActionNo Setter
// 活动期号
func (r *AlibabaMjPresaleSettlementStatisticsRequest) SetActionNo(actionNo int64) error {
    r.actionNo = actionNo
    r.Set("action_no", actionNo)
    return nil
}

// ActionNo Getter
func (r AlibabaMjPresaleSettlementStatisticsRequest) GetActionNo() int64 {
    return r.actionNo
}
// StoreNo Setter
// 外部门店编码
func (r *AlibabaMjPresaleSettlementStatisticsRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMjPresaleSettlementStatisticsRequest) GetStoreNo() string {
    return r.storeNo
}
