package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
呼叫运力 API请求
alibaba.mj.oc.calldispatcher

定时达呼叫运力接口
*/
type AlibabaMjOcCalldispatcherRequest struct {
    model.Params
    // 入参
    callDispatcherDTO   *CallDispatcherDto
}

// 初始化AlibabaMjOcCalldispatcherRequest对象
func NewAlibabaMjOcCalldispatcherRequest() *AlibabaMjOcCalldispatcherRequest{
    return &AlibabaMjOcCalldispatcherRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjOcCalldispatcherRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.calldispatcher"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjOcCalldispatcherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallDispatcherDTO Setter
// 入参
func (r *AlibabaMjOcCalldispatcherRequest) SetCallDispatcherDTO(callDispatcherDTO *CallDispatcherDto) error {
    r.callDispatcherDTO = callDispatcherDTO
    r.Set("call_dispatcher_d_t_o", callDispatcherDTO)
    return nil
}

// CallDispatcherDTO Getter
func (r AlibabaMjOcCalldispatcherRequest) GetCallDispatcherDTO() *CallDispatcherDto {
    return r.callDispatcherDTO
}
