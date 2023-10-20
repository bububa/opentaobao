package tmallsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardDeliveryAPIRequest) Reset() {
	r._identifyTaskDeliveryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardDeliveryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardDeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardDeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallServicecenterWorkcardDeliveryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardDeliveryRequest()
	},
}

// GetTmallServicecenterWorkcardDeliveryRequest 从 sync.Pool 获取 TmallServicecenterWorkcardDeliveryAPIRequest
func GetTmallServicecenterWorkcardDeliveryAPIRequest() *TmallServicecenterWorkcardDeliveryAPIRequest {
	return poolTmallServicecenterWorkcardDeliveryAPIRequest.Get().(*TmallServicecenterWorkcardDeliveryAPIRequest)
}

// ReleaseTmallServicecenterWorkcardDeliveryAPIRequest 将 TmallServicecenterWorkcardDeliveryAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardDeliveryAPIRequest(v *TmallServicecenterWorkcardDeliveryAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardDeliveryAPIRequest.Put(v)
}
