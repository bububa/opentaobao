package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密取消激活回调接口 API请求
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

// 初始化TaobaoAppleCardActiveCancelNotifyRequest对象
func NewTaobaoAppleCardActiveCancelNotifyRequest() *TaobaoAppleCardActiveCancelNotifyRequest{
    return &TaobaoAppleCardActiveCancelNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppleCardActiveCancelNotifyRequest) GetApiMethodName() string {
    return "taobao.apple.card.active.cancel.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppleCardActiveCancelNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 原商户申请激活时的订单号
func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r TaobaoAppleCardActiveCancelNotifyRequest) GetOrderNo() string {
    return r.orderNo
}
// ResultCode Setter
// 结果，000：整单取消激活成功  04： 取消激活失败（包括全部失败和部分失败，此时需以detail为准）
func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

// ResultCode Getter
func (r TaobaoAppleCardActiveCancelNotifyRequest) GetResultCode() string {
    return r.resultCode
}
// GatewayOrderNo Setter
// 网关订单号
func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetGatewayOrderNo(gatewayOrderNo string) error {
    r.gatewayOrderNo = gatewayOrderNo
    r.Set("gateway_order_no", gatewayOrderNo)
    return nil
}

// GatewayOrderNo Getter
func (r TaobaoAppleCardActiveCancelNotifyRequest) GetGatewayOrderNo() string {
    return r.gatewayOrderNo
}
// ResultMsg Setter
// 描述
func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

// ResultMsg Getter
func (r TaobaoAppleCardActiveCancelNotifyRequest) GetResultMsg() string {
    return r.resultMsg
}
// AppleCardCancels Setter
// 卡信息
func (r *TaobaoAppleCardActiveCancelNotifyRequest) SetAppleCardCancels(appleCardCancels []AppleCardCancelDto) error {
    r.appleCardCancels = appleCardCancels
    r.Set("apple_card_cancels", appleCardCancels)
    return nil
}

// AppleCardCancels Getter
func (r TaobaoAppleCardActiveCancelNotifyRequest) GetAppleCardCancels() []AppleCardCancelDto {
    return r.appleCardCancels
}
