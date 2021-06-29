package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据交易号获取物流宝订单 API请求
taobao.wlb.tradeorder.get

根据交易类型和交易id查询物流宝订单详情
*/
type TaobaoWlbTradeorderGetRequest struct {
    model.Params
    // 子交易号
    subTradeId   string
    // 指定交易类型的交易号
    tradeId   string
    // 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
    tradeType   string
}

// 初始化TaobaoWlbTradeorderGetRequest对象
func NewTaobaoWlbTradeorderGetRequest() *TaobaoWlbTradeorderGetRequest{
    return &TaobaoWlbTradeorderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWlbTradeorderGetRequest) GetApiMethodName() string {
    return "taobao.wlb.tradeorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWlbTradeorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubTradeId Setter
// 子交易号
func (r *TaobaoWlbTradeorderGetRequest) SetSubTradeId(subTradeId string) error {
    r.subTradeId = subTradeId
    r.Set("sub_trade_id", subTradeId)
    return nil
}

// SubTradeId Getter
func (r TaobaoWlbTradeorderGetRequest) GetSubTradeId() string {
    return r.subTradeId
}
// TradeId Setter
// 指定交易类型的交易号
func (r *TaobaoWlbTradeorderGetRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

// TradeId Getter
func (r TaobaoWlbTradeorderGetRequest) GetTradeId() string {
    return r.tradeId
}
// TradeType Setter
// 交易类型: TAOBAO--淘宝交易 OTHER_TRADE--其它交易
func (r *TaobaoWlbTradeorderGetRequest) SetTradeType(tradeType string) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoWlbTradeorderGetRequest) GetTradeType() string {
    return r.tradeType
}
