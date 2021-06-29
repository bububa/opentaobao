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
    _actionNo   int64
    // 外部门店编码
    _storeNo   string
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
func (r *AlibabaMjPresaleSettlementStatisticsRequest) SetActionNo(_actionNo int64) error {
    r._actionNo = _actionNo
    r.Set("action_no", _actionNo)
    return nil
}

// ActionNo Getter
func (r AlibabaMjPresaleSettlementStatisticsRequest) GetActionNo() int64 {
    return r._actionNo
}
// StoreNo Setter
// 外部门店编码
func (r *AlibabaMjPresaleSettlementStatisticsRequest) SetStoreNo(_storeNo string) error {
    r._storeNo = _storeNo
    r.Set("store_no", _storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMjPresaleSettlementStatisticsRequest) GetStoreNo() string {
    return r._storeNo
}
