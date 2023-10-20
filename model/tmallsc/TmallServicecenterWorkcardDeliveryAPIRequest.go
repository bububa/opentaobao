package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcarddeliveryAPIRequest 开始配送工单 API请求
// tmall.servicecenter.workcard.delivery
//
// 服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
type TmallservicecenterworkcarddeliveryAPIRequest struct {
	model.Params
	// 工单配送请求参数
	_identifyTaskDeliveryRequest *IdentifyTaskDeliveryRequest
}

// NewTmallservicecenterworkcarddeliveryRequest 初始化TmallservicecenterworkcarddeliveryAPIRequest对象
func NewTmallservicecenterworkcarddeliveryRequest() *TmallservicecenterworkcarddeliveryAPIRequest {
	return &TmallservicecenterworkcarddeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcarddeliveryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcarddeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcarddeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifyTaskDeliveryRequest is IdentifyTaskDeliveryRequest Setter
// 工单配送请求参数
func (r *TmallservicecenterworkcarddeliveryAPIRequest) SetIdentifyTaskDeliveryRequest(_identifyTaskDeliveryRequest *IdentifyTaskDeliveryRequest) error {
	r._identifyTaskDeliveryRequest = _identifyTaskDeliveryRequest
	r.Set("identify_task_delivery_request", _identifyTaskDeliveryRequest)
	return nil
}

// GetIdentifyTaskDeliveryRequest IdentifyTaskDeliveryRequest Getter
func (r TmallservicecenterworkcarddeliveryAPIRequest) GetIdentifyTaskDeliveryRequest() *IdentifyTaskDeliveryRequest {
	return r._identifyTaskDeliveryRequest
}
