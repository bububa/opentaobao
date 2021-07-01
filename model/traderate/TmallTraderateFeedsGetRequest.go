package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询子订单对应的评价、追评以及语义标签 API请求
tmall.traderate.feeds.get

通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
*/
type TmallTraderateFeedsGetAPIRequest struct {
    model.Params
    // 交易子订单ID
    _childTradeId   int64
}

// 初始化TmallTraderateFeedsGetAPIRequest对象
func NewTmallTraderateFeedsGetRequest() *TmallTraderateFeedsGetAPIRequest{
    return &TmallTraderateFeedsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraderateFeedsGetAPIRequest) GetApiMethodName() string {
    return "tmall.traderate.feeds.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraderateFeedsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChildTradeId Setter
// 交易子订单ID
func (r *TmallTraderateFeedsGetAPIRequest) SetChildTradeId(_childTradeId int64) error {
    r._childTradeId = _childTradeId
    r.Set("child_trade_id", _childTradeId)
    return nil
}

// ChildTradeId Getter
func (r TmallTraderateFeedsGetAPIRequest) GetChildTradeId() int64 {
    return r._childTradeId
}
