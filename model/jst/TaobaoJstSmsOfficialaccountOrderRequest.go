package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔订购公众号 APIRequest
taobao.jst.sms.officialaccount.order

聚石塔订购公众号接口
*/
type TaobaoJstSmsOfficialaccountOrderRequest struct {
    model.Params

    // 聚石塔公众号订购
    orderOfficialAccountRequest   *OrderOfficialAccountRequest 

}

func NewTaobaoJstSmsOfficialaccountOrderRequest() *TaobaoJstSmsOfficialaccountOrderRequest{
    return &TaobaoJstSmsOfficialaccountOrderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsOfficialaccountOrderRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.order"
}

func (r TaobaoJstSmsOfficialaccountOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsOfficialaccountOrderRequest) SetOrderOfficialAccountRequest(orderOfficialAccountRequest *OrderOfficialAccountRequest) error {
    r.orderOfficialAccountRequest = orderOfficialAccountRequest
    r.Set("order_official_account_request", orderOfficialAccountRequest)
    return nil
}

func (r TaobaoJstSmsOfficialaccountOrderRequest) GetOrderOfficialAccountRequest() *OrderOfficialAccountRequest {
    return r.orderOfficialAccountRequest
}

