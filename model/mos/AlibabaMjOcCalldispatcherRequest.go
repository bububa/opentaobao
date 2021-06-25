package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
呼叫运力 APIRequest
alibaba.mj.oc.calldispatcher

定时达呼叫运力接口
*/
type AlibabaMjOcCalldispatcherRequest struct {
    model.Params

    // 入参
    callDispatcherDTO   *CallDispatcherDto 

}

func NewAlibabaMjOcCalldispatcherRequest() *AlibabaMjOcCalldispatcherRequest{
    return &AlibabaMjOcCalldispatcherRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjOcCalldispatcherRequest) GetApiMethodName() string {
    return "alibaba.mj.oc.calldispatcher"
}

func (r AlibabaMjOcCalldispatcherRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjOcCalldispatcherRequest) SetCallDispatcherDTO(callDispatcherDTO *CallDispatcherDto) error {
    r.callDispatcherDTO = callDispatcherDTO
    r.Set("call_dispatcher_d_t_o", callDispatcherDTO)
    return nil
}

func (r AlibabaMjOcCalldispatcherRequest) GetCallDispatcherDTO() *CallDispatcherDto {
    return r.callDispatcherDTO
}

