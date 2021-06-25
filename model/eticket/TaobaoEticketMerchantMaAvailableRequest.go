package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证核销前校验接口 APIRequest
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

func NewTaobaoEticketMerchantMaAvailableRequest() *TaobaoEticketMerchantMaAvailableRequest{
    return &TaobaoEticketMerchantMaAvailableRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.available"
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoEticketMerchantMaAvailableRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetBizType() int64 {
    return r.bizType
}

func (r *TaobaoEticketMerchantMaAvailableRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetCode() string {
    return r.code
}

func (r *TaobaoEticketMerchantMaAvailableRequest) SetConsumeNum(consumeNum int64) error {
    r.consumeNum = consumeNum
    r.Set("consume_num", consumeNum)
    return nil
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetConsumeNum() int64 {
    return r.consumeNum
}

func (r *TaobaoEticketMerchantMaAvailableRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoEticketMerchantMaAvailableRequest) SetPosId(posId string) error {
    r.posId = posId
    r.Set("pos_id", posId)
    return nil
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetPosId() string {
    return r.posId
}

func (r *TaobaoEticketMerchantMaAvailableRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetSerialNum() string {
    return r.serialNum
}

func (r *TaobaoEticketMerchantMaAvailableRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoEticketMerchantMaAvailableRequest) GetToken() string {
    return r.token
}

