package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家确认收货 APIRequest
tmall.exchange.returngoods.agree

卖家确认收货
*/
type TmallExchangeReturngoodsAgreeRequest struct {
    model.Params

    // 换货单号ID
    disputeId   int64 

    // 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
    fields   []string 

}

func NewTmallExchangeReturngoodsAgreeRequest() *TmallExchangeReturngoodsAgreeRequest{
    return &TmallExchangeReturngoodsAgreeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeReturngoodsAgreeRequest) GetApiMethodName() string {
    return "tmall.exchange.returngoods.agree"
}

func (r TmallExchangeReturngoodsAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeReturngoodsAgreeRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeReturngoodsAgreeRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeReturngoodsAgreeRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeReturngoodsAgreeRequest) GetFields() []string {
    return r.fields
}

