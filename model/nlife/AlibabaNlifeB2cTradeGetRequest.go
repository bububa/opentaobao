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
    tradeNo   string
    // 外部订单号，和trade_no不能同时为空
    outTradeNo   string
    // 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
    storeId   string
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
func (r *AlibabaNlifeB2cTradeGetRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaNlifeB2cTradeGetRequest) GetTradeNo() string {
    return r.tradeNo
}
// OutTradeNo Setter
// 外部订单号，和trade_no不能同时为空
func (r *AlibabaNlifeB2cTradeGetRequest) SetOutTradeNo(outTradeNo string) error {
    r.outTradeNo = outTradeNo
    r.Set("out_trade_no", outTradeNo)
    return nil
}

// OutTradeNo Getter
func (r AlibabaNlifeB2cTradeGetRequest) GetOutTradeNo() string {
    return r.outTradeNo
}
// StoreId Setter
// 零售+门店ID，如果传递的是outTradeNola，那么这个是必传的
func (r *AlibabaNlifeB2cTradeGetRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaNlifeB2cTradeGetRequest) GetStoreId() string {
    return r.storeId
}
