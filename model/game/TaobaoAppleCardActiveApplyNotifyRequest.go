package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密申请激活回调接口 APIRequest
taobao.apple.card.active.apply.notify

苹果卡密申请激活回调接口
*/
type TaobaoAppleCardActiveApplyNotifyRequest struct {
    model.Params

    // 卡列表
    appleCards   []AppleCardDto 

    // 网关订单号
    gatewayOrderNo   string 

    // 描述
    resultMsg   string 

    // 商户唯一订单号
    orderNo   string 

    // 结果，000：成功，其他皆为错误 04： 订单处理失败(商户可退款，其他不可退款)
    resultCode   string 

}

func NewTaobaoAppleCardActiveApplyNotifyRequest() *TaobaoAppleCardActiveApplyNotifyRequest{
    return &TaobaoAppleCardActiveApplyNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAppleCardActiveApplyNotifyRequest) GetApiMethodName() string {
    return "taobao.apple.card.active.apply.notify"
}

func (r TaobaoAppleCardActiveApplyNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAppleCardActiveApplyNotifyRequest) SetAppleCards(appleCards []AppleCardDto) error {
    r.appleCards = appleCards
    r.Set("apple_cards", appleCards)
    return nil
}

func (r TaobaoAppleCardActiveApplyNotifyRequest) GetAppleCards() []AppleCardDto {
    return r.appleCards
}

func (r *TaobaoAppleCardActiveApplyNotifyRequest) SetGatewayOrderNo(gatewayOrderNo string) error {
    r.gatewayOrderNo = gatewayOrderNo
    r.Set("gateway_order_no", gatewayOrderNo)
    return nil
}

func (r TaobaoAppleCardActiveApplyNotifyRequest) GetGatewayOrderNo() string {
    return r.gatewayOrderNo
}

func (r *TaobaoAppleCardActiveApplyNotifyRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

func (r TaobaoAppleCardActiveApplyNotifyRequest) GetResultMsg() string {
    return r.resultMsg
}

func (r *TaobaoAppleCardActiveApplyNotifyRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r TaobaoAppleCardActiveApplyNotifyRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *TaobaoAppleCardActiveApplyNotifyRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

func (r TaobaoAppleCardActiveApplyNotifyRequest) GetResultCode() string {
    return r.resultCode
}

