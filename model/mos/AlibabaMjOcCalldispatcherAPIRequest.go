package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjoccalldispatcherAPIRequest 呼叫运力 API请求
// alibaba.mj.oc.calldispatcher
//
// 定时达呼叫运力接口
type AlibabamjoccalldispatcherAPIRequest struct {
	model.Params
	// 入参
	_callDispatcherDTO *CallDispatcherDto
}

// NewAlibabamjoccalldispatcherRequest 初始化AlibabamjoccalldispatcherAPIRequest对象
func NewAlibabamjoccalldispatcherRequest() *AlibabamjoccalldispatcherAPIRequest {
	return &AlibabamjoccalldispatcherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjoccalldispatcherAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.calldispatcher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjoccalldispatcherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjoccalldispatcherAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallDispatcherDTO is CallDispatcherDTO Setter
// 入参
func (r *AlibabamjoccalldispatcherAPIRequest) SetCallDispatcherDTO(_callDispatcherDTO *CallDispatcherDto) error {
	r._callDispatcherDTO = _callDispatcherDTO
	r.Set("call_dispatcher_d_t_o", _callDispatcherDTO)
	return nil
}

// GetCallDispatcherDTO CallDispatcherDTO Getter
func (r AlibabamjoccalldispatcherAPIRequest) GetCallDispatcherDTO() *CallDispatcherDto {
	return r._callDispatcherDTO
}
