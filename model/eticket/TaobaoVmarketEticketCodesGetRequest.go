package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证码列表查询 APIRequest
taobao.vmarket.eticket.codes.get

查询某个订单的所有码的列表
*/
type TaobaoVmarketEticketCodesGetRequest struct {
    model.Params

    // 订单号
    orderId   int64 

    // 码商ID
    codemerchantId   int64 

}

func NewTaobaoVmarketEticketCodesGetRequest() *TaobaoVmarketEticketCodesGetRequest{
    return &TaobaoVmarketEticketCodesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketCodesGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.codes.get"
}

func (r TaobaoVmarketEticketCodesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketCodesGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVmarketEticketCodesGetRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoVmarketEticketCodesGetRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

func (r TaobaoVmarketEticketCodesGetRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}

