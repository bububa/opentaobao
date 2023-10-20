package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmstaskcreateAPIRequest 聚石塔短信任务创建接口 API请求
// taobao.jst.sms.task.create
//
// 聚石塔短信的任务创建接口，用于创建数字短信、公众号短信、权益短信的AB测试任务。
type TaobaojstsmstaskcreateAPIRequest struct {
	model.Params
	// 创建任务的入参
	_paramCreateSmsTaskRequest *CreateSmsTaskRequest
}

// NewTaobaojstsmstaskcreateRequest 初始化TaobaojstsmstaskcreateAPIRequest对象
func NewTaobaojstsmstaskcreateRequest() *TaobaojstsmstaskcreateAPIRequest {
	return &TaobaojstsmstaskcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmstaskcreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.task.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmstaskcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmstaskcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCreateSmsTaskRequest is ParamCreateSmsTaskRequest Setter
// 创建任务的入参
func (r *TaobaojstsmstaskcreateAPIRequest) SetParamCreateSmsTaskRequest(_paramCreateSmsTaskRequest *CreateSmsTaskRequest) error {
	r._paramCreateSmsTaskRequest = _paramCreateSmsTaskRequest
	r.Set("param_create_sms_task_request", _paramCreateSmsTaskRequest)
	return nil
}

// GetParamCreateSmsTaskRequest ParamCreateSmsTaskRequest Getter
func (r TaobaojstsmstaskcreateAPIRequest) GetParamCreateSmsTaskRequest() *CreateSmsTaskRequest {
	return r._paramCreateSmsTaskRequest
}
