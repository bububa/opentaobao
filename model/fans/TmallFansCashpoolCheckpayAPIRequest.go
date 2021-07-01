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
type TmallFansCashpoolCheckpayAPIRequest struct {
    model.Params
    // 资金池列表
    _cashPoolList   []int64
}

// 初始化TmallFansCashpoolCheckpayAPIRequest对象
func NewTmallFansCashpoolCheckpayRequest() *TmallFansCashpoolCheckpayAPIRequest{
    return &TmallFansCashpoolCheckpayAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallFansCashpoolCheckpayAPIRequest) GetApiMethodName() string {
    return "tmall.fans.cashpool.checkpay"
}

// IRequest interface 方法, 获取API参数
func (r TmallFansCashpoolCheckpayAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CashPoolList Setter
// 资金池列表
func (r *TmallFansCashpoolCheckpayAPIRequest) SetCashPoolList(_cashPoolList []int64) error {
    r._cashPoolList = _cashPoolList
    r.Set("cash_pool_list", _cashPoolList)
    return nil
}

// CashPoolList Getter
func (r TmallFansCashpoolCheckpayAPIRequest) GetCashPoolList() []int64 {
    return r._cashPoolList
}
