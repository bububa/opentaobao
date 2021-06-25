package traderate

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询子订单对应的评价、追评以及语义标签 APIRequest
tmall.traderate.feeds.get

通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
*/
type TmallTraderateFeedsGetRequest struct {
    model.Params

    // 交易子订单ID
    childTradeId   int64 

}

func NewTmallTraderateFeedsGetRequest() *TmallTraderateFeedsGetRequest{
    return &TmallTraderateFeedsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraderateFeedsGetRequest) GetApiMethodName() string {
    return "tmall.traderate.feeds.get"
}

func (r TmallTraderateFeedsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraderateFeedsGetRequest) SetChildTradeId(childTradeId int64) error {
    r.childTradeId = childTradeId
    r.Set("child_trade_id", childTradeId)
    return nil
}

func (r TmallTraderateFeedsGetRequest) GetChildTradeId() int64 {
    return r.childTradeId
}

