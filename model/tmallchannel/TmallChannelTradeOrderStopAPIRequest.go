package tmallchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商停止发货 API请求
tmall.channel.trade.order.stop

供应商停止发货
*/
type TmallChannelTradeOrderStopAPIRequest struct {
    model.Params
    // 主采购单号
    _mainPurchaseOrderNo   int64
    // 幂等单号
    _requestNo   string
}

// 初始化TmallChannelTradeOrderStopAPIRequest对象
func NewTmallChannelTradeOrderStopRequest() *TmallChannelTradeOrderStopAPIRequest{
    return &TmallChannelTradeOrderStopAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallChannelTradeOrderStopAPIRequest) GetApiMethodName() string {
    return "tmall.channel.trade.order.stop"
}

// IRequest interface 方法, 获取API参数
func (r TmallChannelTradeOrderStopAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainPurchaseOrderNo Setter
// 主采购单号
func (r *TmallChannelTradeOrderStopAPIRequest) SetMainPurchaseOrderNo(_mainPurchaseOrderNo int64) error {
    r._mainPurchaseOrderNo = _mainPurchaseOrderNo
    r.Set("main_purchase_order_no", _mainPurchaseOrderNo)
    return nil
}

// MainPurchaseOrderNo Getter
func (r TmallChannelTradeOrderStopAPIRequest) GetMainPurchaseOrderNo() int64 {
    return r._mainPurchaseOrderNo
}
// RequestNo Setter
// 幂等单号
func (r *TmallChannelTradeOrderStopAPIRequest) SetRequestNo(_requestNo string) error {
    r._requestNo = _requestNo
    r.Set("request_no", _requestNo)
    return nil
}

// RequestNo Getter
func (r TmallChannelTradeOrderStopAPIRequest) GetRequestNo() string {
    return r._requestNo
}
