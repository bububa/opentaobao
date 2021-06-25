package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务核销 APIRequest
tmall.service.code.consume

天猫服务平台－服务核销
*/
type TmallServiceCodeConsumeRequest struct {
    model.Params

    // 核销码
    consumeCode   string 

    // 核销帐号
    operatorNick   string 

    // 门店id
    shopId   string 

}

func NewTmallServiceCodeConsumeRequest() *TmallServiceCodeConsumeRequest{
    return &TmallServiceCodeConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServiceCodeConsumeRequest) GetApiMethodName() string {
    return "tmall.service.code.consume"
}

func (r TmallServiceCodeConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServiceCodeConsumeRequest) SetConsumeCode(consumeCode string) error {
    r.consumeCode = consumeCode
    r.Set("consume_code", consumeCode)
    return nil
}

func (r TmallServiceCodeConsumeRequest) GetConsumeCode() string {
    return r.consumeCode
}

func (r *TmallServiceCodeConsumeRequest) SetOperatorNick(operatorNick string) error {
    r.operatorNick = operatorNick
    r.Set("operator_nick", operatorNick)
    return nil
}

func (r TmallServiceCodeConsumeRequest) GetOperatorNick() string {
    return r.operatorNick
}

func (r *TmallServiceCodeConsumeRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r TmallServiceCodeConsumeRequest) GetShopId() string {
    return r.shopId
}

