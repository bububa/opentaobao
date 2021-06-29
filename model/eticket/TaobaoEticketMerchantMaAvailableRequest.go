package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销前校验接口 API请求
taobao.eticket.merchant.ma.available

商家验码之前的调用接口，用来判断是否可以进行核销操作
*/
type TaobaoEticketMerchantMaAvailableRequest struct {
    model.Params
    // 业务类型
    bizType   int64
    // 需要被核销的码
    code   string
    // 核销份数
    consumeNum   int64
    // 业务id（订单号）
    outerId   string
    // 机具编号
    posId   string
    // 核销序列号，需要保证唯一
    serialNum   string
    // 需要跟发码通知获取到的参数一致
    token   string
}

// 初始化TaobaoEticketMerchantMaAvailableRequest对象
func NewTaobaoEticketMerchantMaAvailableRequest() *TaobaoEticketMerchantMaAvailableRequest{
    return &TaobaoEticketMerchantMaAvailableRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoEticketMerchantMaAvailableRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.available"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoEticketMerchantMaAvailableRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizType Setter
// 业务类型
func (r *TaobaoEticketMerchantMaAvailableRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

// BizType Getter
func (r TaobaoEticketMerchantMaAvailableRequest) GetBizType() int64 {
    return r.bizType
}
// Code Setter
// 需要被核销的码
func (r *TaobaoEticketMerchantMaAvailableRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r TaobaoEticketMerchantMaAvailableRequest) GetCode() string {
    return r.code
}
// ConsumeNum Setter
// 核销份数
func (r *TaobaoEticketMerchantMaAvailableRequest) SetConsumeNum(consumeNum int64) error {
    r.consumeNum = consumeNum
    r.Set("consume_num", consumeNum)
    return nil
}

// ConsumeNum Getter
func (r TaobaoEticketMerchantMaAvailableRequest) GetConsumeNum() int64 {
    return r.consumeNum
}
// OuterId Setter
// 业务id（订单号）
func (r *TaobaoEticketMerchantMaAvailableRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoEticketMerchantMaAvailableRequest) GetOuterId() string {
    return r.outerId
}
// PosId Setter
// 机具编号
func (r *TaobaoEticketMerchantMaAvailableRequest) SetPosId(posId string) error {
    r.posId = posId
    r.Set("pos_id", posId)
    return nil
}

// PosId Getter
func (r TaobaoEticketMerchantMaAvailableRequest) GetPosId() string {
    return r.posId
}
// SerialNum Setter
// 核销序列号，需要保证唯一
func (r *TaobaoEticketMerchantMaAvailableRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

// SerialNum Getter
func (r TaobaoEticketMerchantMaAvailableRequest) GetSerialNum() string {
    return r.serialNum
}
// Token Setter
// 需要跟发码通知获取到的参数一致
func (r *TaobaoEticketMerchantMaAvailableRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoEticketMerchantMaAvailableRequest) GetToken() string {
    return r.token
}
