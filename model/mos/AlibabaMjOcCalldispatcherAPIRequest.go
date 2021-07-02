package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcCalldispatcherAPIRequest 呼叫运力 API请求
// alibaba.mj.oc.calldispatcher
//
// 定时达呼叫运力接口
type AlibabaMjOcCalldispatcherAPIRequest struct {
	model.Params
	// 入参
	_callDispatcherDTO *CallDispatcherDto
}

// NewAlibabaMjOcCalldispatcherRequest 初始化AlibabaMjOcCalldispatcherAPIRequest对象
func NewAlibabaMjOcCalldispatcherRequest() *AlibabaMjOcCalldispatcherAPIRequest {
	return &AlibabaMjOcCalldispatcherAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcCalldispatcherAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.calldispatcher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcCalldispatcherAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCallDispatcherDTO is CallDispatcherDTO Setter
// 入参
func (r *AlibabaMjOcCalldispatcherAPIRequest) SetCallDispatcherDTO(_callDispatcherDTO *CallDispatcherDto) error {
	r._callDispatcherDTO = _callDispatcherDTO
	r.Set("call_dispatcher_d_t_o", _callDispatcherDTO)
	return nil
}

// GetCallDispatcherDTO CallDispatcherDTO Getter
func (r AlibabaMjOcCalldispatcherAPIRequest) GetCallDispatcherDTO() *CallDispatcherDto {
	return r._callDispatcherDTO
}
