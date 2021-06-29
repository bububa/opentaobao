package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通退款订单查询 API请求
alibaba.lst.trade.refund.order.get

零售通退款订单查询
*/
type AlibabaLstTradeRefundOrderGetRequest struct {
    model.Params
    // 退款单id
    _refundId   string
    // 主订单id
    _mainOrderId   int64
}

// 初始化AlibabaLstTradeRefundOrderGetRequest对象
func NewAlibabaLstTradeRefundOrderGetRequest() *AlibabaLstTradeRefundOrderGetRequest{
    return &AlibabaLstTradeRefundOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeRefundOrderGetRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.refund.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeRefundOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款单id
func (r *AlibabaLstTradeRefundOrderGetRequest) SetRefundId(_refundId string) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r AlibabaLstTradeRefundOrderGetRequest) GetRefundId() string {
    return r._refundId
}
// MainOrderId Setter
// 主订单id
func (r *AlibabaLstTradeRefundOrderGetRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r AlibabaLstTradeRefundOrderGetRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
