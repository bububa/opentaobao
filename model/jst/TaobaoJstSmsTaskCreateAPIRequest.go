package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsTaskCreateAPIRequest 聚石塔短信任务创建接口 API请求
// taobao.jst.sms.task.create
//
// 聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
type TaobaoJstSmsTaskCreateAPIRequest struct {
	model.Params
	// 创建任务的入参
	_paramCreateSmsTaskRequest *CreateSmsTaskRequest
}

// NewTaobaoJstSmsTaskCreateRequest 初始化TaobaoJstSmsTaskCreateAPIRequest对象
func NewTaobaoJstSmsTaskCreateRequest() *TaobaoJstSmsTaskCreateAPIRequest {
	return &TaobaoJstSmsTaskCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsTaskCreateAPIRequest) Reset() {
	r._paramCreateSmsTaskRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTaskCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.task.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTaskCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsTaskCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCreateSmsTaskRequest is ParamCreateSmsTaskRequest Setter
// 创建任务的入参
func (r *TaobaoJstSmsTaskCreateAPIRequest) SetParamCreateSmsTaskRequest(_paramCreateSmsTaskRequest *CreateSmsTaskRequest) error {
	r._paramCreateSmsTaskRequest = _paramCreateSmsTaskRequest
	r.Set("param_create_sms_task_request", _paramCreateSmsTaskRequest)
	return nil
}

// GetParamCreateSmsTaskRequest ParamCreateSmsTaskRequest Getter
func (r TaobaoJstSmsTaskCreateAPIRequest) GetParamCreateSmsTaskRequest() *CreateSmsTaskRequest {
	return r._paramCreateSmsTaskRequest
}

var poolTaobaoJstSmsTaskCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsTaskCreateRequest()
	},
}

// GetTaobaoJstSmsTaskCreateRequest 从 sync.Pool 获取 TaobaoJstSmsTaskCreateAPIRequest
func GetTaobaoJstSmsTaskCreateAPIRequest() *TaobaoJstSmsTaskCreateAPIRequest {
	return poolTaobaoJstSmsTaskCreateAPIRequest.Get().(*TaobaoJstSmsTaskCreateAPIRequest)
}

// ReleaseTaobaoJstSmsTaskCreateAPIRequest 将 TaobaoJstSmsTaskCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsTaskCreateAPIRequest(v *TaobaoJstSmsTaskCreateAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsTaskCreateAPIRequest.Put(v)
}
