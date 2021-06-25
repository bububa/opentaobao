package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密取消激活回调接口 APIRequest
taobao.apple.card.active.cancel.notify

苹果卡密取消激活回调接口
*/
type TaobaoAppleCardActiveCancelNotifyRequest struct {
    model.Params

    // 原商户申请激活时的订单号
    orderNo   string 

    // 结果，000：整单取消激活成功  04： 取消激活失败（包括全部失败和部分失败，此时需以detail为准）
    resultCode   string 

    // 网关订单号
    gatewayOrderNo   string 

    // 描述
    resultMsg   string 

    // 卡信息
    appleCardCancels   []AppleCardCancelDto 

}

func NewTaobaoAppleCardActiveCancelNotifyRequest() *TaobaoAppleCardActiveCancelNotifyRequest{
    return &TaobaoAppleCardActiveCancelNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAppleCardActiveCancelNotifyRequest) GetApiMethodName() string {
    return "taobao.apple.card.active.cancel.notify"
}

func (r TaobaoAppleCardActiveCancelNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r TaobaoAppleCardActiveCancelNotifyRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

func (r TaobaoAppleCardActiveCancelNotifyRequest) GetResultCode() string {
    return r.resultCode
}

func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetGatewayOrderNo(gatewayOrderNo string) error {
    r.gatewayOrderNo = gatewayOrderNo
    r.Set("gateway_order_no", gatewayOrderNo)
    return nil
}

func (r TaobaoAppleCardActiveCancelNotifyRequest) GetGatewayOrderNo() string {
    return r.gatewayOrderNo
}

func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

func (r TaobaoAppleCardActiveCancelNotifyRequest) GetResultMsg() string {
    return r.resultMsg
}

func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetAppleCardCancels(appleCardCancels []AppleCardCancelDto) error {
    r.appleCardCancels = appleCardCancels
    r.Set("apple_card_cancels", appleCardCancels)
    return nil
}

func (r TaobaoAppleCardActiveCancelNotifyRequest) GetAppleCardCancels() []AppleCardCancelDto {
    return r.appleCardCancels
}

