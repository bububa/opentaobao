package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
呼叫运力 APIResponse
alibaba.mj.oc.calldispatcher

定时达呼叫运力接口
*/
type AlibabaMjOcCalldispatcherAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_oc_calldispatcher_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultDO `json:"result,omitempty" xml:"