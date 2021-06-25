package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动权益发放接口 APIRequest
taobao.promotion.benefit.activity.send

活动权益发放接口，用于卖家针对活动进行权益发放
*/
type TaobaoPromotionBenefitActivitySendRequest struct {
    model.Params

    // 单个权益发放请求
    sendRequest   *BenefitSingleSendRequest 

    // 非混淆的接收者id
    receiverId   int64 

    // 混淆的接收者nick
    nick   string 

    // 非混淆的接收者nick
    platNick   string 

    // 混淆的接收者id
    mixReceiverId   string 

}

func NewTaobaoPromotionBenefitActivitySendRequest() *TaobaoPromotionBenefitActivitySendRequest{
    return &TaobaoPromotionBenefitActivitySendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionBenefitActivitySendRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.send"
}

func (r TaobaoPromotionBenefitActivitySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionBenefitActivitySendRequest) SetSendRequest(sendRequest *BenefitSingleSendRequest) error {
    r.sendRequest = sendRequest
    r.Set("send_request", sendRequest)
    return nil
}

func (r TaobaoPromotionBenefitActivitySendRequest) GetSendRequest() *BenefitSingleSendRequest {
    return r.sendRequest
}

func (r *TaobaoPromotionBenefitActivitySendRequest) SetReceiverId(receiverId int64) error {
    r.receiverId = receiverId
    r.Set("receiver_id", receiverId)
    return nil
}

func (r TaobaoPromotionBenefitActivitySendRequest) GetReceiverId() int64 {
    return r.receiverId
}

func (r *TaobaoPromotionBenefitActivitySendRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoPromotionBenefitActivitySendRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoPromotionBenefitActivitySendRequest) SetPlatNick(platNick string) error {
    r.platNick = platNick
    r.Set("plat_nick", platNick)
    return nil
}

func (r TaobaoPromotionBenefitActivitySendRequest) GetPlatNick() string {
    return r.platNick
}

func (r *TaobaoPromotionBenefitActivitySendRequest) SetMixReceiverId(mixReceiverId string) error {
    r.mixReceiverId = mixReceiverId
    r.Set("mix_receiver_id", mixReceiverId)
    return nil
}

func (r TaobaoPromotionBenefitActivitySendRequest) GetMixReceiverId() string {
    return r.mixReceiverId
}

