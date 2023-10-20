package jym

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryTaskswitchSaveAPIRequest 行业信息系统开关 API请求
// alibaba.jym.industry.taskswitch.save
//
// VMOS回调交易猫行业信息系统
type AlibabaJymIndustryTaskswitchSaveAPIRequest struct {
	model.Params
	// 开关参数
	_taskSwitchReqDto *TaskSwitchReqDto
}

// NewAlibabaJymIndustryTaskswitchSaveRequest 初始化AlibabaJymIndustryTaskswitchSaveAPIRequest对象
func NewAlibabaJymIndustryTaskswitchSaveRequest() *AlibabaJymIndustryTaskswitchSaveAPIRequest {
	return &AlibabaJymIndustryTaskswitchSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymIndustryTaskswitchSaveAPIRequest) Reset() {
	r._taskSwitchReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymIndustryTaskswitchSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.taskswitch.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymIndustryTaskswitchSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymIndustryTaskswitchSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaskSwitchReqDto is TaskSwitchReqDto Setter
// 开关参数
func (r *AlibabaJymIndustryTaskswitchSaveAPIRequest) SetTaskSwitchReqDto(_taskSwitchReqDto *TaskSwitchReqDto) error {
	r._taskSwitchReqDto = _taskSwitchReqDto
	r.Set("task_switch_req_dto", _taskSwitchReqDto)
	return nil
}

// GetTaskSwitchReqDto TaskSwitchReqDto Getter
func (r AlibabaJymIndustryTaskswitchSaveAPIRequest) GetTaskSwitchReqDto() *TaskSwitchReqDto {
	return r._taskSwitchReqDto
}

var poolAlibabaJymIndustryTaskswitchSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymIndustryTaskswitchSaveRequest()
	},
}

// GetAlibabaJymIndustryTaskswitchSaveRequest 从 sync.Pool 获取 AlibabaJymIndustryTaskswitchSaveAPIRequest
func GetAlibabaJymIndustryTaskswitchSaveAPIRequest() *AlibabaJymIndustryTaskswitchSaveAPIRequest {
	return poolAlibabaJymIndustryTaskswitchSaveAPIRequest.Get().(*AlibabaJymIndustryTaskswitchSaveAPIRequest)
}

// ReleaseAlibabaJymIndustryTaskswitchSaveAPIRequest 将 AlibabaJymIndustryTaskswitchSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymIndustryTaskswitchSaveAPIRequest(v *AlibabaJymIndustryTaskswitchSaveAPIRequest) {
	v.Reset()
	poolAlibabaJymIndustryTaskswitchSaveAPIRequest.Put(v)
}
