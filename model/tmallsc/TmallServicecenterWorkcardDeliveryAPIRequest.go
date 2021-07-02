package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardDeliveryAPIRequest 开始配送工单 API请求
// tmall.servicecenter.workcard.delivery
//
// 服务商调用该接口通知天猫服务平台服务商工人已开始配送工单
type TmallServicecenterWorkcardDeliveryAPIRequest struct {
	model.Params
	// 工单配送请求参数
	_identifyTaskDeliveryRequest *IdentifyTaskDeliveryRequest
}

// NewTmallServicecenterWorkcardDeliveryRequest 初始化TmallServicecenterWorkcardDeliveryAPIRequest对象
func NewTmallServicecenterWorkcardDeliveryRequest() *TmallServicecenterWorkcardDeliveryAPIRequest {
	return &TmallServicecenterWorkcardDeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardDeliveryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardDeliveryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIdentifyTaskDeliveryRequest is IdentifyTaskDeliveryRequest Setter
// 工单配送请求参数
func (r *TmallServicecenterWorkcardDeliveryAPIRequest) SetIdentifyTaskDeliveryRequest(_identifyTaskDeliveryRequest *IdentifyTaskDeliveryRequest) error {
	r._identifyTaskDeliveryRequest = _identifyTaskDeliveryRequest
	r.Set("identify_task_delivery_request", _identifyTaskDeliveryRequest)
	return nil
}

// GetIdentifyTaskDeliveryRequest IdentifyTaskDeliveryRequest Getter
func (r TmallServicecenterWorkcardDeliveryAPIRequest) GetIdentifyTaskDeliveryRequest() *IdentifyTaskDeliveryRequest {
	return r._identifyTaskDeliveryRequest
}
