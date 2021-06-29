package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
检查资金池付款状态 API请求
tmall.fans.cashpool.checkpay

检查资金池付款状态
*/
type TmallFansCashpoolCheckpayRequest struct {
    model.Params
    // 资金池列表
    cashPoolList   []int64
}

// 初始化TmallFansCashpoolCheckpayRequest对象
func NewTmallFansCashpoolCheckpayRequest() *TmallFansCashpoolCheckpayRequest{
    return &TmallFansCashpoolCheckpayRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFansCashpoolCheckpayRequest) GetApiMethodName() string {
    return "tmall.fans.cashpool.checkpay"
}

// IRequest interface 方法, 获取API参数
func (r TmallFansCashpoolCheckpayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CashPoolList Setter
// 资金池列表
func (r *TmallFansCashpoolCheckpayRequest) SetCashPoolList(cashPoolList []int64) error {
    r.cashPoolList = cashPoolList
    r.Set("cash_pool_list", cashPoolList)
    return nil
}

// CashPoolList Getter
func (r TmallFansCashpoolCheckpayRequest) GetCashPoolList() []int64 {
    return r.cashPoolList
}
