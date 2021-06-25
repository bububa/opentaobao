package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 APIRequest
taobao.eticket.merchant.ma.reverse

电子凭证平台冲正接口
*/
type TaobaoEticketMerchantMaReverseRequest struct {
    model.Params

    // 业务类型
    bizType   int64 

    // 码值
    code   string 

    // 业务id（订单号）
    outerId   string 

    // 机具编号，如果核销时有则必传
    posId   string 

    // 冲正份数，需要与核销份数一致
    reverseNum   int64 

    // 需要冲正的核销序列号
    serialNum   string 

    // 需要跟发码通知获取到的参数一致
    token   string 

}

func NewTaobaoEticketMerchantMaReverseRequest() *TaobaoEticketMerchantMaReverseRequest{
    return &TaobaoEticketMerchantMaReverseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoEticketMerchantMaReverseRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.reverse"
}

func (r TaobaoEticketMerchantMaReverseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoEticketMerchantMaReverseRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoEticketMerchantMaReverseRequest) GetBizType() int64 {
    return r.bizType
}

func (r *TaobaoEticketMerchantMaReverseRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r TaobaoEticketMerchantMaReverseRequest) GetCode() string {
    return r.code
}

func (r *TaobaoEticketMerchantMaReverseRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoEticketMerchantMaReverseRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoEticketMerchantMaReverseRequest) SetPosId(posId string) error {
    r.posId = posId
    r.Set("pos_id", posId)
    return nil
}

func (r TaobaoEticketMerchantMaReverseRequest) GetPosId() string {
    return r.posId
}

func (r *TaobaoEticketMerchantMaReverseRequest) SetReverseNum(reverseNum int64) error {
    r.reverseNum = reverseNum
    r.Set("reverse_num", reverseNum)
    return nil
}

func (r TaobaoEticketMerchantMaReverseRequest) GetReverseNum() int64 {
    return r.reverseNum
}

func (r *TaobaoEticketMerchantMaReverseRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

func (r TaobaoEticketMerchantMaReverseRequest) GetSerialNum() string {
    return r.serialNum
}

func (r *TaobaoEticketMerchantMaReverseRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoEticketMerchantMaReverseRequest) GetToken() string {
    return r.token
}

