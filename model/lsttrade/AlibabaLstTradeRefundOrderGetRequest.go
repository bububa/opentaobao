package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通退款订单查询 APIRequest
alibaba.lst.trade.refund.order.get

零售通退款订单查询
*/
type AlibabaLstTradeRefundOrderGetRequest struct {
    model.Params

    // 退款单id
    refundId   string 

    // 主订单id
    mainOrderId   int64 

}

func NewAlibabaLstTradeRefundOrderGetRequest() *AlibabaLstTradeRefundOrderGetRequest{
    return &AlibabaLstTradeRefundOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeRefundOrderGetRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.refund.order.get"
}

func (r AlibabaLstTradeRefundOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeRefundOrderGetRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r AlibabaLstTradeRefundOrderGetRequest) GetRefundId() string {
    return r.refundId
}

func (r *AlibabaLstTradeRefundOrderGetRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r AlibabaLstTradeRefundOrderGetRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

