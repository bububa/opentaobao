package fans

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
检查资金池付款状态 APIRequest
tmall.fans.cashpool.checkpay

检查资金池付款状态
*/
type TmallFansCashpoolCheckpayRequest struct {
    model.Params

    // 资金池列表
    cashPoolList   []int64 

}

func NewTmallFansCashpoolCheckpayRequest() *TmallFansCashpoolCheckpayRequest{
    return &TmallFansCashpoolCheckpayRequest{
        Params: model.NewParams(),
    }
}

func (r TmallFansCashpoolCheckpayRequest) GetApiMethodName() string {
    return "tmall.fans.cashpool.checkpay"
}

func (r TmallFansCashpoolCheckpayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallFansCashpoolCheckpayRequest) SetCashPoolList(cashPoolList []int64) error {
    r.cashPoolList = cashPoolList
    r.Set("cash_pool_list", cashPoolList)
    return nil
}

func (r TmallFansCashpoolCheckpayRequest) GetCashPoolList() []int64 {
    return r.cashPoolList
}

