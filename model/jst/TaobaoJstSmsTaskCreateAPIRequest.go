package jst

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsTaskCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.task.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsTaskCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamCreateSmsTaskRequest Setter
// 创建任务的入参
func (r *TaobaoJstSmsTaskCreateAPIRequest) SetParamCreateSmsTaskRequest(_paramCreateSmsTaskRequest *CreateSmsTaskRequest) error {
	r._paramCreateSmsTaskRequest = _paramCreateSmsTaskRequest
	r.Set("param_create_sms_task_request", _paramCreateSmsTaskRequest)
	return nil
}

// Get ParamCreateSmsTaskRequest Getter
func (r TaobaoJstSmsTaskCreateAPIRequest) GetParamCreateSmsTaskRequest() *CreateSmsTaskRequest {
	return r._paramCreateSmsTaskRequest
}
