package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密申请激活回调接口 API请求
taobao.apple.card.active.apply.notify

苹果卡密申请激活回调接口
*/
type TaobaoAppleCardActiveApplyNotifyAPIRequest struct {
    model.Params
    // 卡列表
    _appleCards   []AppleCardDto
    // 网关订单号
    _gatewayOrderNo   string
    // 描述
    _resultMsg   string
    // 商户唯一订单号
    _orderNo   string
    // 结果，000：成功，其他皆为错误 04： 订单处理失败(商户可退款，其他不可退款)
    _resultCode   string
}

// 初始化TaobaoAppleCardActiveApplyNotifyAPIRequest对象
func NewTaobaoAppleCardActiveApplyNotifyRequest() *TaobaoAppleCardActiveApplyNotifyAPIRequest{
    return &TaobaoAppleCardActiveApplyNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetApiMethodName() string {
    return "taobao.apple.card.active.apply.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppleCards Setter
// 卡列表
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetAppleCards(_appleCards []AppleCardDto) error {
    r._appleCards = _appleCards
    r.Set("apple_cards", _appleCards)
    return nil
}

// AppleCards Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetAppleCards() []AppleCardDto {
    return r._appleCards
}
// GatewayOrderNo Setter
// 网关订单号
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetGatewayOrderNo(_gatewayOrderNo string) error {
    r._gatewayOrderNo = _gatewayOrderNo
    r.Set("gateway_order_no", _gatewayOrderNo)
    return nil
}

// GatewayOrderNo Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetGatewayOrderNo() string {
    return r._gatewayOrderNo
}
// ResultMsg Setter
// 描述
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
    r._resultMsg = _resultMsg
    r.Set("result_msg", _resultMsg)
    return nil
}

// ResultMsg Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetResultMsg() string {
    return r._resultMsg
}
// OrderNo Setter
// 商户唯一订单号
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetOrderNo() string {
    return r._orderNo
}
// ResultCode Setter
// 结果，000：成功，其他皆为错误 04： 订单处理失败(商户可退款，其他不可退款)
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetResultCode(_resultCode string) error {
    r._resultCode = _resultCode
    r.Set("result_code", _resultCode)
    return nil
}

// ResultCode Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetResultCode() string {
    return r._resultCode
}
