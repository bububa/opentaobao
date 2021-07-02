package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketManageNotifyAPIRequest 主动发起通知接口 API请求
// taobao.vmarket.eticket.manage.notify
//
// 外部合作商家主动发起通知接口
type TaobaoVmarketEticketManageNotifyAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
	// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
	_codemerchantId int64
	// 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
	_notifyMethod string
}

// NewTaobaoVmarketEticketManageNotifyRequest 初始化TaobaoVmarketEticketManageNotifyAPIRequest对象
func NewTaobaoVmarketEticketManageNotifyRequest() *TaobaoVmarketEticketManageNotifyAPIRequest {
	return &TaobaoVmarketEticketManageNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketManageNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.manage.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketManageNotifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *TaobaoVmarketEticketManageNotifyAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketManageNotifyAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaoVmarketEticketManageNotifyAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaoVmarketEticketManageNotifyAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}

// SetNotifyMethod is NotifyMethod Setter
// 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
func (r *TaobaoVmarketEticketManageNotifyAPIRequest) SetNotifyMethod(_notifyMethod string) error {
	r._notifyMethod = _notifyMethod
	r.Set("notify_method", _notifyMethod)
	return nil
}

// GetNotifyMethod NotifyMethod Getter
func (r TaobaoVmarketEticketManageNotifyAPIRequest) GetNotifyMethod() string {
	return r._notifyMethod
}
