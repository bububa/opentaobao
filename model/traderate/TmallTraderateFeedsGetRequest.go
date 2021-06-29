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
type TmallTraderateFeedsGetRequest struct {
    model.Params
    // 交易子订单ID
    _childTradeId   int64
}

// 初始化TmallTraderateFeedsGetRequest对象
func NewTmallTraderateFeedsGetRequest() *TmallTraderateFeedsGetRequest{
    return &TmallTraderateFeedsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraderateFeedsGetRequest) GetApiMethodName() string {
    return "tmall.traderate.feeds.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraderateFeedsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChildTradeId Setter
// 交易子订单ID
func (r *TmallTraderateFeedsGetRequest) SetChildTradeId(_childTradeId int64) error {
    r._childTradeId = _childTradeId
    r.Set("child_trade_id", _childTradeId)
    return nil
}

// ChildTradeId Getter
func (r TmallTraderateFeedsGetRequest) GetChildTradeId() int64 {
    return r._childTradeId
}
