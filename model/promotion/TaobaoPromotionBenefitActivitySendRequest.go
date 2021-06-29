package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
活动权益发放接口 API请求
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

// 初始化TaobaoPromotionBenefitActivitySendRequest对象
func NewTaobaoPromotionBenefitActivitySendRequest() *TaobaoPromotionBenefitActivitySendRequest{
    return &TaobaoPromotionBenefitActivitySendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivitySendRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivitySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SendRequest Setter
// 单个权益发放请求
func (r *TaobaoPromotionBenefitActivitySendRequest) SetSendRequest(sendRequest *BenefitSingleSendRequest) error {
    r.sendRequest = sendRequest
    r.Set("send_request", sendRequest)
    return nil
}

// SendRequest Getter
func (r TaobaoPromotionBenefitActivitySendRequest) GetSendRequest() *BenefitSingleSendRequest {
    return r.sendRequest
}
// ReceiverId Setter
// 非混淆的接收者id
func (r *TaobaoPromotionBenefitActivitySendRequest) SetReceiverId(receiverId int64) error {
    r.receiverId = receiverId
    r.Set("receiver_id", receiverId)
    return nil
}

// ReceiverId Getter
func (r TaobaoPromotionBenefitActivitySendRequest) GetReceiverId() int64 {
    return r.receiverId
}
// Nick Setter
// 混淆的接收者nick
func (r *TaobaoPromotionBenefitActivitySendRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoPromotionBenefitActivitySendRequest) GetNick() string {
    return r.nick
}
// PlatNick Setter
// 非混淆的接收者nick
func (r *TaobaoPromotionBenefitActivitySendRequest) SetPlatNick(platNick string) error {
    r.platNick = platNick
    r.Set("plat_nick", platNick)
    return nil
}

// PlatNick Getter
func (r TaobaoPromotionBenefitActivitySendRequest) GetPlatNick() string {
    return r.platNick
}
// MixReceiverId Setter
// 混淆的接收者id
func (r *TaobaoPromotionBenefitActivitySendRequest) SetMixReceiverId(mixReceiverId string) error {
    r.mixReceiverId = mixReceiverId
    r.Set("mix_receiver_id", mixReceiverId)
    return nil
}

// MixReceiverId Getter
func (r TaobaoPromotionBenefitActivitySendRequest) GetMixReceiverId() string {
    return r.mixReceiverId
}
