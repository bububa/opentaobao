package larkiot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkIotOrderConfirmorderAPIRequest iot渠道卖品落单 API请求
// taobao.lark.iot.order.confirmorder
//
// 云智对接无人超市，接收无人超市订单信息
type TaobaoLarkIotOrderConfirmorderAPIRequest struct {
	model.Params
	// 渠道编码
	_channelCode string
	// 影院内码
	_cinemaLinkId string
	// 外部订单号
	_outGoodsOrderId string
	// 工作站id
	_workstationId string
	// 工作站名称
	_workstationName string
	// 支付方式
	_paymentList string
	// 优惠列表
	_promotionList string
	// 卖品列表
	_goodsList string
	// 手机号
	_mobile string
	// 管理员
	_operatorUserId string
}

// NewTaobaoLarkIotOrderConfirmorderRequest 初始化TaobaoLarkIotOrderConfirmorderAPIRequest对象
func NewTaobaoLarkIotOrderConfirmorderRequest() *TaobaoLarkIotOrderConfirmorderAPIRequest {
	return &TaobaoLarkIotOrderConfirmorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetApiMethodName() string {
	return "taobao.lark.iot.order.confirmorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelCode is ChannelCode Setter
// 渠道编码
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetChannelCode(_channelCode string) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetChannelCode() string {
	return r._channelCode
}

// SetCinemaLinkId is CinemaLinkId Setter
// 影院内码
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetCinemaLinkId(_cinemaLinkId string) error {
	r._cinemaLinkId = _cinemaLinkId
	r.Set("cinema_link_id", _cinemaLinkId)
	return nil
}

// GetCinemaLinkId CinemaLinkId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetCinemaLinkId() string {
	return r._cinemaLinkId
}

// SetOutGoodsOrderId is OutGoodsOrderId Setter
// 外部订单号
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetOutGoodsOrderId(_outGoodsOrderId string) error {
	r._outGoodsOrderId = _outGoodsOrderId
	r.Set("out_goods_order_id", _outGoodsOrderId)
	return nil
}

// GetOutGoodsOrderId OutGoodsOrderId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetOutGoodsOrderId() string {
	return r._outGoodsOrderId
}

// SetWorkstationId is WorkstationId Setter
// 工作站id
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetWorkstationId(_workstationId string) error {
	r._workstationId = _workstationId
	r.Set("workstation_id", _workstationId)
	return nil
}

// GetWorkstationId WorkstationId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetWorkstationId() string {
	return r._workstationId
}

// SetWorkstationName is WorkstationName Setter
// 工作站名称
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetWorkstationName(_workstationName string) error {
	r._workstationName = _workstationName
	r.Set("workstation_name", _workstationName)
	return nil
}

// GetWorkstationName WorkstationName Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetWorkstationName() string {
	return r._workstationName
}

// SetPaymentList is PaymentList Setter
// 支付方式
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetPaymentList(_paymentList string) error {
	r._paymentList = _paymentList
	r.Set("payment_list", _paymentList)
	return nil
}

// GetPaymentList PaymentList Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetPaymentList() string {
	return r._paymentList
}

// SetPromotionList is PromotionList Setter
// 优惠列表
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetPromotionList(_promotionList string) error {
	r._promotionList = _promotionList
	r.Set("promotion_list", _promotionList)
	return nil
}

// GetPromotionList PromotionList Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetPromotionList() string {
	return r._promotionList
}

// SetGoodsList is GoodsList Setter
// 卖品列表
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetGoodsList(_goodsList string) error {
	r._goodsList = _goodsList
	r.Set("goods_list", _goodsList)
	return nil
}

// GetGoodsList GoodsList Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetGoodsList() string {
	return r._goodsList
}

// SetMobile is Mobile Setter
// 手机号
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetMobile() string {
	return r._mobile
}

// SetOperatorUserId is OperatorUserId Setter
// 管理员
func (r *TaobaoLarkIotOrderConfirmorderAPIRequest) SetOperatorUserId(_operatorUserId string) error {
	r._operatorUserId = _operatorUserId
	r.Set("operator_user_id", _operatorUserId)
	return nil
}

// GetOperatorUserId OperatorUserId Getter
func (r TaobaoLarkIotOrderConfirmorderAPIRequest) GetOperatorUserId() string {
	return r._operatorUserId
}
