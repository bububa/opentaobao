package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家确认收货 API请求
tmall.exchange.returngoods.agree

卖家确认收货
*/
type TmallExchangeReturngoodsAgreeRequest struct {
    model.Params
    // 换货单号ID
    _disputeId   int64
    // 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
    _fields   []string
}

// 初始化TmallExchangeReturngoodsAgreeRequest对象
func NewTmallExchangeReturngoodsAgreeRequest() *TmallExchangeReturngoodsAgreeRequest{
    return &TmallExchangeReturngoodsAgreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeReturngoodsAgreeRequest) GetApiMethodName() string {
    return "tmall.exchange.returngoods.agree"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeReturngoodsAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeReturngoodsAgreeRequest) SetDisputeId(_disputeId int64) error {
    r._disputeId = _disputeId
    r.Set("dispute_id", _disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeReturngoodsAgreeRequest) GetDisputeId() int64 {
    return r._disputeId
}
// Fields Setter
// 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
func (r *TmallExchangeReturngoodsAgreeRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TmallExchangeReturngoodsAgreeRequest) GetFields() []string {
    return r._fields
}
