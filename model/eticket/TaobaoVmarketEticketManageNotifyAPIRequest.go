package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketmanagenotifyAPIRequest 主动发起通知接口 API请求
// taobao.vmarket.eticket.manage.notify
//
// 外部合作商家主动发起通知接口
type TaobaovmarketeticketmanagenotifyAPIRequest struct {
	model.Params
	// 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
	_notifyMethod string
	// 订单编号
	_orderId int64
	// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
	_codemerchantId int64
}

// NewTaobaovmarketeticketmanagenotifyRequest 初始化TaobaovmarketeticketmanagenotifyAPIRequest对象
func NewTaobaovmarketeticketmanagenotifyRequest() *TaobaovmarketeticketmanagenotifyAPIRequest {
	return &TaobaovmarketeticketmanagenotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketmanagenotifyAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.manage.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketmanagenotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketmanagenotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNotifyMethod is NotifyMethod Setter
// 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
func (r *TaobaovmarketeticketmanagenotifyAPIRequest) SetNotifyMethod(_notifyMethod string) error {
	r._notifyMethod = _notifyMethod
	r.Set("notify_method", _notifyMethod)
	return nil
}

// GetNotifyMethod NotifyMethod Getter
func (r TaobaovmarketeticketmanagenotifyAPIRequest) GetNotifyMethod() string {
	return r._notifyMethod
}

// SetOrderId is OrderId Setter
// 订单编号
func (r *TaobaovmarketeticketmanagenotifyAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovmarketeticketmanagenotifyAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaovmarketeticketmanagenotifyAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaovmarketeticketmanagenotifyAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
