package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售+平台查询订单 API请求
alibaba.nlife.b2c.trade.get

查询零售+平台创建出来的订单详情
*/
type AlibabaNlifeB2cTradeGetRequest struct {
    model.Params
    // 零售+平台订单号,和out_trade_no不能同时为空
    _tradeNo   string
    // 外部订单号，和trade_no不能同时为空
    _outTradeNo   string
    // 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
    _storeId   string
}

// 初始化AlibabaNlifeB2cTradeGetRequest对象
func NewAlibabaNlifeB2cTradeGetRequest() *AlibabaNlifeB2cTradeGetRequest{
    return &AlibabaNlifeB2cTradeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2cTradeGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2c.trade.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2cTradeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeNo Setter
// 零售+平台订单号,和out_trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeGetRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaNlifeB2cTradeGetRequest) GetTradeNo() string {
    return r._tradeNo
}
// OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeGetRequest) SetOutTradeNo(_outTradeNo string) error {
    r._outTradeNo = _outTradeNo
    r.Set("out_trade_no", _outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r AlibabaNlifeB2cTradeGetRequest) GetOutTradeNo() string {
    return r._outTradeNo
}
// StoreId Setter
// 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
func (r *AlibabaNlifeB2cTradeGetRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradeGetRequest) GetStoreId() string {
    return r._storeId
}
