package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymindustrytaskswitchsaveAPIRequest 行业信息系统开关 API请求
// alibaba.jym.industry.taskswitch.save
//
// VMOS回调交易猫行业信息系统
type AlibabajymindustrytaskswitchsaveAPIRequest struct {
	model.Params
	// 开关参数
	_taskSwitchReqDto *TaskSwitchReqDto
}

// NewAlibabajymindustrytaskswitchsaveRequest 初始化AlibabajymindustrytaskswitchsaveAPIRequest对象
func NewAlibabajymindustrytaskswitchsaveRequest() *AlibabajymindustrytaskswitchsaveAPIRequest {
	return &AlibabajymindustrytaskswitchsaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymindustrytaskswitchsaveAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.taskswitch.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymindustrytaskswitchsaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymindustrytaskswitchsaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskSwitchReqDto is TaskSwitchReqDto Setter
// 开关参数
func (r *AlibabajymindustrytaskswitchsaveAPIRequest) SetTaskSwitchReqDto(_taskSwitchReqDto *TaskSwitchReqDto) error {
	r._taskSwitchReqDto = _taskSwitchReqDto
	r.Set("task_switch_req_dto", _taskSwitchReqDto)
	return nil
}

// GetTaskSwitchReqDto TaskSwitchReqDto Getter
func (r AlibabajymindustrytaskswitchsaveAPIRequest) GetTaskSwitchReqDto() *TaskSwitchReqDto {
	return r._taskSwitchReqDto
}
