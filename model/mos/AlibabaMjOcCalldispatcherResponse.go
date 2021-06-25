package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
呼叫运力 APIResponse
alibaba.mj.oc.calldispatcher

定时达呼叫运力接口
*/
type AlibabaMjOcCalldispatcherAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjOcCalldispatcherResponse `json:"alibaba_mj_oc_calldispatcher_response,omitempty"`
}

type AlibabaMjOcCalldispatcherResponse struct {

    // result
    Result   *ResultDO `json:"result,omitempty"`

}
