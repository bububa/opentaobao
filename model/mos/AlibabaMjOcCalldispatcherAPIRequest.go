package mos

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjOcCalldispatcherAPIRequest) Reset() {
	r._callDispatcherDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcCalldispatcherAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.calldispatcher"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcCalldispatcherAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcCalldispatcherAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaMjOcCalldispatcherAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjOcCalldispatcherRequest()
	},
}

// GetAlibabaMjOcCalldispatcherRequest 从 sync.Pool 获取 AlibabaMjOcCalldispatcherAPIRequest
func GetAlibabaMjOcCalldispatcherAPIRequest() *AlibabaMjOcCalldispatcherAPIRequest {
	return poolAlibabaMjOcCalldispatcherAPIRequest.Get().(*AlibabaMjOcCalldispatcherAPIRequest)
}

// ReleaseAlibabaMjOcCalldispatcherAPIRequest 将 AlibabaMjOcCalldispatcherAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjOcCalldispatcherAPIRequest(v *AlibabaMjOcCalldispatcherAPIRequest) {
	v.Reset()
	poolAlibabaMjOcCalldispatcherAPIRequest.Put(v)
}
