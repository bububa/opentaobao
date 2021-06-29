package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证码列表查询 API请求
taobao.vmarket.eticket.codes.get

查询某个订单的所有码的列表
*/
type TaobaoVmarketEticketCodesGetRequest struct {
    model.Params
    // 订单号
    _orderId   int64
    // 码商ID
    _codemerchantId   int64
}

// 初始化TaobaoVmarketEticketCodesGetRequest对象
func NewTaobaoVmarketEticketCodesGetRequest() *TaobaoVmarketEticketCodesGetRequest{
    return &TaobaoVmarketEticketCodesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketCodesGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.codes.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketCodesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单号
func (r *TaobaoVmarketEticketCodesGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketCodesGetRequest) GetOrderId() int64 {
    return r._orderId
}
// CodemerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketCodesGetRequest) SetCodemerchantId(_codemerchantId int64) error {
    r._codemerchantId = _codemerchantId
    r.Set("codemerchant_id", _codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketCodesGetRequest) GetCodemerchantId() int64 {
    return r._codemerchantId
}
