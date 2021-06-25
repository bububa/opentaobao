package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单优惠账单查询 APIRequest
alibaba.wdk.trade.discount.bill.get

商家查询订单优惠账单
*/
type AlibabaWdkTradeDiscountBillGetRequest struct {
    model.Params

    // 请求参数
    param0   *OrderDiscountBillQueryRequest 

}

func NewAlibabaWdkTradeDiscountBillGetRequest() *AlibabaWdkTradeDiscountBillGetRequest{
    return &AlibabaWdkTradeDiscountBillGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkTradeDiscountBillGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.discount.bill.get"
}

func (r AlibabaWdkTradeDiscountBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkTradeDiscountBillGetRequest) SetParam0(param0 *OrderDiscountBillQueryRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaWdkTradeDiscountBillGetRequest) GetParam0() *OrderDiscountBillQueryRequest {
    return r.param0
}

