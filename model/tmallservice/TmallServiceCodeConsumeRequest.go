package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台服务核销 API请求
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

// 初始化TmallServiceCodeConsumeRequest对象
func NewTmallServiceCodeConsumeRequest() *TmallServiceCodeConsumeRequest{
    return &TmallServiceCodeConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServiceCodeConsumeRequest) GetApiMethodName() string {
    return "tmall.service.code.consume"
}

// IRequest interface 方法, 获取API参数
func (r TmallServiceCodeConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsumeCode Setter
// 核销码
func (r *TmallServiceCodeConsumeRequest) SetConsumeCode(consumeCode string) error {
    r.consumeCode = consumeCode
    r.Set("consume_code", consumeCode)
    return nil
}

// ConsumeCode Getter
func (r TmallServiceCodeConsumeRequest) GetConsumeCode() string {
    return r.consumeCode
}
// OperatorNick Setter
// 核销帐号
func (r *TmallServiceCodeConsumeRequest) SetOperatorNick(operatorNick string) error {
    r.operatorNick = operatorNick
    r.Set("operator_nick", operatorNick)
    return nil
}

// OperatorNick Getter
func (r TmallServiceCodeConsumeRequest) GetOperatorNick() string {
    return r.operatorNick
}
// ShopId Setter
// 门店id
func (r *TmallServiceCodeConsumeRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r TmallServiceCodeConsumeRequest) GetShopId() string {
    return r.shopId
}
