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
type TmallExchangeReturngoodsAgreeAPIRequest struct {
    model.Params
    // 换货单号ID
    _disputeId   int64
    // 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
    _fields   []string
}

// 初始化TmallExchangeReturngoodsAgreeAPIRequest对象
func NewTmallExchangeReturngoodsAgreeRequest() *TmallExchangeReturngoodsAgreeAPIRequest{
    return &TmallExchangeReturngoodsAgreeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetApiMethodName() string {
    return "tmall.exchange.returngoods.agree"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeReturngoodsAgreeAPIRequest) SetDisputeId(_disputeId int64) error {
    r._disputeId = _disputeId
    r.Set("dispute_id", _disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetDisputeId() int64 {
    return r._disputeId
}
// Fields Setter
// 返回字段。目前支持dispute_id（换货单号ID）,bizorder_id（正向交易单号ID）, modified（订单修改时间）, status（当前换货状态）
func (r *TmallExchangeReturngoodsAgreeAPIRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TmallExchangeReturngoodsAgreeAPIRequest) GetFields() []string {
    return r._fields
}
