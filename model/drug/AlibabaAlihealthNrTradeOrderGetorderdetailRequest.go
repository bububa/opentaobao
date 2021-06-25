package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单id获取单条订单详情 APIRequest
alibaba.alihealth.nr.trade.order.getorderdetail

阿里健康O2O，获取订单详情，修复组合商品价格精度问题
*/
type AlibabaAlihealthNrTradeOrderGetorderdetailRequest struct {
    model.Params

    // 淘宝订单ID
    orderId   int64 

}

func NewAlibabaAlihealthNrTradeOrderGetorderdetailRequest() *AlibabaAlihealthNrTradeOrderGetorderdetailRequest{
    return &AlibabaAlihealthNrTradeOrderGetorderdetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthNrTradeOrderGetorderdetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.trade.order.getorderdetail"
}

func (r AlibabaAlihealthNrTradeOrderGetorderdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthNrTradeOrderGetorderdetailRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthNrTradeOrderGetorderdetailRequest) GetOrderId() int64 {
    return r.orderId
}

