package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressdeliverycutnotifyAPIRequest TMS配拦截结果回告 API请求
// taobao.logistics.express.delivery.cut.notify
//
// TMS配拦截结果回告
type TaobaologisticsexpressdeliverycutnotifyAPIRequest struct {
	model.Params
	// 请求
	_notifyCutOffDeliveryProcessStatusRequest *NotifyCutOffDeliveryProcessStatusRequest
}

// NewTaobaologisticsexpressdeliverycutnotifyRequest 初始化TaobaologisticsexpressdeliverycutnotifyAPIRequest对象
func NewTaobaologisticsexpressdeliverycutnotifyRequest() *TaobaologisticsexpressdeliverycutnotifyAPIRequest {
	return &TaobaologisticsexpressdeliverycutnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressdeliverycutnotifyAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.delivery.cut.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressdeliverycutnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressdeliverycutnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNotifyCutOffDeliveryProcessStatusRequest is NotifyCutOffDeliveryProcessStatusRequest Setter
// 请求
func (r *TaobaologisticsexpressdeliverycutnotifyAPIRequest) SetNotifyCutOffDeliveryProcessStatusRequest(_notifyCutOffDeliveryProcessStatusRequest *NotifyCutOffDeliveryProcessStatusRequest) error {
	r._notifyCutOffDeliveryProcessStatusRequest = _notifyCutOffDeliveryProcessStatusRequest
	r.Set("notify_cut_off_delivery_process_status_request", _notifyCutOffDeliveryProcessStatusRequest)
	return nil
}

// GetNotifyCutOffDeliveryProcessStatusRequest NotifyCutOffDeliveryProcessStatusRequest Getter
func (r TaobaologisticsexpressdeliverycutnotifyAPIRequest) GetNotifyCutOffDeliveryProcessStatusRequest() *NotifyCutOffDeliveryProcessStatusRequest {
	return r._notifyCutOffDeliveryProcessStatusRequest
}
