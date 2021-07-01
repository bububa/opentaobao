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
type TaobaoPromotionBenefitActivitySendAPIRequest struct {
    model.Params
    // 单个权益发放请求
    _sendRequest   *BenefitSingleSendRequest
    // 非混淆的接收者id
    _receiverId   int64
    // 混淆的接收者nick
    _nick   string
    // 非混淆的接收者nick
    _platNick   string
    // 混淆的接收者id
    _mixReceiverId   string
}

// 初始化TaobaoPromotionBenefitActivitySendAPIRequest对象
func NewTaobaoPromotionBenefitActivitySendRequest() *TaobaoPromotionBenefitActivitySendAPIRequest{
    return &TaobaoPromotionBenefitActivitySendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetApiMethodName() string {
    return "taobao.promotion.benefit.activity.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SendRequest Setter
// 单个权益发放请求
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetSendRequest(_sendRequest *BenefitSingleSendRequest) error {
    r._sendRequest = _sendRequest
    r.Set("send_request", _sendRequest)
    return nil
}

// SendRequest Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetSendRequest() *BenefitSingleSendRequest {
    return r._sendRequest
}
// ReceiverId Setter
// 非混淆的接收者id
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetReceiverId(_receiverId int64) error {
    r._receiverId = _receiverId
    r.Set("receiver_id", _receiverId)
    return nil
}

// ReceiverId Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetReceiverId() int64 {
    return r._receiverId
}
// Nick Setter
// 混淆的接收者nick
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetNick() string {
    return r._nick
}
// PlatNick Setter
// 非混淆的接收者nick
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetPlatNick(_platNick string) error {
    r._platNick = _platNick
    r.Set("plat_nick", _platNick)
    return nil
}

// PlatNick Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetPlatNick() string {
    return r._platNick
}
// MixReceiverId Setter
// 混淆的接收者id
func (r *TaobaoPromotionBenefitActivitySendAPIRequest) SetMixReceiverId(_mixReceiverId string) error {
    r._mixReceiverId = _mixReceiverId
    r.Set("mix_receiver_id", _mixReceiverId)
    return nil
}

// MixReceiverId Getter
func (r TaobaoPromotionBenefitActivitySendAPIRequest) GetMixReceiverId() string {
    return r._mixReceiverId
}
