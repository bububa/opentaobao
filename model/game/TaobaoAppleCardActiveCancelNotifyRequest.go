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
type TaobaoAppleCardActiveCancelNotifyAPIRequest struct {
    model.Params
    // 原商户申请激活时的订单号
    _orderNo   string
    // 结果，000：整单取消激活成功  04： 取消激活失败（包括全部失败和部分失败，此时需以detail为准）
    _resultCode   string
    // 网关订单号
    _gatewayOrderNo   string
    // 描述
    _resultMsg   string
    // 卡信息
    _appleCardCancels   []AppleCardCancelDTO
}

// 初始化TaobaoAppleCardActiveCancelNotifyAPIRequest对象
func NewTaobaoAppleCardActiveCancelNotifyRequest() *TaobaoAppleCardActiveCancelNotifyAPIRequest{
    return &TaobaoAppleCardActiveCancelNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetApiMethodName() string {
    return "taobao.apple.card.active.cancel.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 原商户申请激活时的订单号
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetOrderNo() string {
    return r._orderNo
}
// ResultCode Setter
// 结果，000：整单取消激活成功  04： 取消激活失败（包括全部失败和部分失败，此时需以detail为准）
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetResultCode(_resultCode string) error {
    r._resultCode = _resultCode
    r.Set("result_code", _resultCode)
    return nil
}

// ResultCode Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetResultCode() string {
    return r._resultCode
}
// GatewayOrderNo Setter
// 网关订单号
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetGatewayOrderNo(_gatewayOrderNo string) error {
    r._gatewayOrderNo = _gatewayOrderNo
    r.Set("gateway_order_no", _gatewayOrderNo)
    return nil
}

// GatewayOrderNo Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetGatewayOrderNo() string {
    return r._gatewayOrderNo
}
// ResultMsg Setter
// 描述
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
    r._resultMsg = _resultMsg
    r.Set("result_msg", _resultMsg)
    return nil
}

// ResultMsg Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetResultMsg() string {
    return r._resultMsg
}
// AppleCardCancels Setter
// 卡信息
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetAppleCardCancels(_appleCardCancels []AppleCardCancelDTO) error {
    r._appleCardCancels = _appleCardCancels
    r.Set("apple_card_cancels", _appleCardCancels)
    return nil
}

// AppleCardCancels Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetAppleCardCancels() []AppleCardCancelDTO {
    return r._appleCardCancels
}
