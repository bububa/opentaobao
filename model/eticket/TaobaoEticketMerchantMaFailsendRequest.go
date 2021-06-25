package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码商发码失败回调接口 APIRequest
taobao.eticket.merchant.ma.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoEticketMerchantMaFailsendRequest struct {
    model.Params

    // 业务id（订单号）
    outerId   string 

    // 错误原因码
    subErrCode   string 

    // 错误码描述
    subErrMsg   string 

    // 需要与发码通知获取的值一致
    token   string 

    // 业务类型
    bizType   int64 

}

func NewTaobaoEticketMerchantMaFailsendRequest() *TaobaoEticketMerchantMaFailsendRequest{
    return &TaobaoEticketMerchantMaFailsendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoEticketMerchantMaFailsendRequest) GetApiMethodName() string {
    return "taobao.eticket.merchant.ma.failsend"
}

func (r TaobaoEticketMerchantMaFailsendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoEticketMerchantMaFailsendRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoEticketMerchantMaFailsendRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoEticketMerchantMaFailsendRequest) SetSubErrCode(subErrCode string) error {
    r.subErrCode = subErrCode
    r.Set("sub_err_code", subErrCode)
    return nil
}

func (r TaobaoEticketMerchantMaFailsendRequest) GetSubErrCode() string {
    return r.subErrCode
}

func (r *TaobaoEticketMerchantMaFailsendRequest) SetSubErrMsg(subErrMsg string) error {
    r.subErrMsg = subErrMsg
    r.Set("sub_err_msg", subErrMsg)
    return nil
}

func (r TaobaoEticketMerchantMaFailsendRequest) GetSubErrMsg() string {
    return r.subErrMsg
}

func (r *TaobaoEticketMerchantMaFailsendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoEticketMerchantMaFailsendRequest) GetToken() string {
    return r.token
}

func (r *TaobaoEticketMerchantMaFailsendRequest) SetBizType(bizType int64) error {
    r.bizType = bizType
    r.Set("biz_type", bizType)
    return nil
}

func (r TaobaoEticketMerchantMaFailsendRequest) GetBizType() int64 {
    return r.bizType
}

