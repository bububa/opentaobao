package larkiot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
iot渠道卖品落单 API请求
taobao.lark.iot.order.confirmorder

云智对接无人超市，接收无人超市订单信息
*/
type TaobaoLarkIotOrderConfirmorderAPIRequest struct {
    model.Params
    // 渠道编码
    _channelCode   string
    // 影院内码
    _cinemaLinkId   string
    // 外部订单号
    _outGoodsOrderId   string
    // 工作站id
    _workstationId   string
    // 工作站名称
    _workstationName   string
    // 支付方式
    _paymentList   string
    // 优惠列表
    _promotionList   string
    // 卖品列表
    _goodsList   string
    // 手机号
    _mobile   string
    // 管理员
    _operatorUserId   string
}

// 初始化TaobaoLarkIotOrderConfirmorderAPIRequest对象
func NewTaobaoLarkIotOrderConfirmorderRequest() *TaobaoLarkIotOrderConfirmorderAPIRequest{
    return &TaobaoLarkIotOrderConfirmorderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetApiMethodName() string {
    return "taobao.lark.iot.order.confirmorder"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetChannelCode(_channelCode string) error {
    r._channelCode = _channelCode
    r.Set("channel_code", _channelCode)
    return nil
}

// ChannelCode Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetChannelCode() string {
    return r._channelCode
}
// CinemaLinkId Setter
// 影院内码
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetCinemaLinkId(_cinemaLinkId string) error {
    r._cinemaLinkId = _cinemaLinkId
    r.Set("cinema_link_id", _cinemaLinkId)
    return nil
}

// CinemaLinkId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetCinemaLinkId() string {
    return r._cinemaLinkId
}
// OutGoodsOrderId Setter
// 外部订单号
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetOutGoodsOrderId(_outGoodsOrderId string) error {
    r._outGoodsOrderId = _outGoodsOrderId
    r.Set("out_goods_order_id", _outGoodsOrderId)
    return nil
}

// OutGoodsOrderId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetOutGoodsOrderId() string {
    return r._outGoodsOrderId
}
// WorkstationId Setter
// 工作站id
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetWorkstationId(_workstationId string) error {
    r._workstationId = _workstationId
    r.Set("workstation_id", _workstationId)
    return nil
}

// WorkstationId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetWorkstationId() string {
    return r._workstationId
}
// WorkstationName Setter
// 工作站名称
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetWorkstationName(_workstationName string) error {
    r._workstationName = _workstationName
    r.Set("workstation_name", _workstationName)
    return nil
}

// WorkstationName Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetWorkstationName() string {
    return r._workstationName
}
// PaymentList Setter
// 支付方式
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetPaymentList(_paymentList string) error {
    r._paymentList = _paymentList
    r.Set("payment_list", _paymentList)
    return nil
}

// PaymentList Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetPaymentList() string {
    return r._paymentList
}
// PromotionList Setter
// 优惠列表
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetPromotionList(_promotionList string) error {
    r._promotionList = _promotionList
    r.Set("promotion_list", _promotionList)
    return nil
}

// PromotionList Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetPromotionList() string {
    return r._promotionList
}
// GoodsList Setter
// 卖品列表
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetGoodsList(_goodsList string) error {
    r._goodsList = _goodsList
    r.Set("goods_list", _goodsList)
    return nil
}

// GoodsList Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetGoodsList() string {
    return r._goodsList
}
// Mobile Setter
// 手机号
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetMobile() string {
    return r._mobile
}
// OperatorUserId Setter
// 管理员
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetOperatorUserId(_operatorUserId string) error {
    r._operatorUserId = _operatorUserId
    r.Set("operator_user_id", _operatorUserId)
    return nil
}

// OperatorUserId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetOperatorUserId() string {
    return r._operatorUserId
}
