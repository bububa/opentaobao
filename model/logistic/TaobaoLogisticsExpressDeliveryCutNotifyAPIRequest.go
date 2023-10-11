package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest TMS配拦截结果回告 API请求
// taobao.logistics.express.delivery.cut.notify
//
// TMS配拦截结果回告
type TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest struct {
	model.Params
	// 请求
	_notifyCutOffDeliveryProcessStatusRequest *NotifyCutOffDeliveryProcessStatusRequest
}

// NewTaobaoLogisticsExpressDeliveryCutNotifyRequest 初始化TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest对象
func NewTaobaoLogisticsExpressDeliveryCutNotifyRequest() *TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest {
	return &TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.delivery.cut.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNotifyCutOffDeliveryProcessStatusRequest is NotifyCutOffDeliveryProcessStatusRequest Setter
// 请求
func (r *TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest) SetNotifyCutOffDeliveryProcessStatusRequest(_notifyCutOffDeliveryProcessStatusRequest *NotifyCutOffDeliveryProcessStatusRequest) error {
	r._notifyCutOffDeliveryProcessStatusRequest = _notifyCutOffDeliveryProcessStatusRequest
	r.Set("notify_cut_off_delivery_process_status_request", _notifyCutOffDeliveryProcessStatusRequest)
	return nil
}

// GetNotifyCutOffDeliveryProcessStatusRequest NotifyCutOffDeliveryProcessStatusRequest Getter
func (r TaobaoLogisticsExpressDeliveryCutNotifyAPIRequest) GetNotifyCutOffDeliveryProcessStatusRequest() *NotifyCutOffDeliveryProcessStatusRequest {
	return r._notifyCutOffDeliveryProcessStatusRequest
}
