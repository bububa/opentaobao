package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
pos故障报警 APIResponse
alibaba.mos.pos.alarm

故障报警
*/
type AlibabaMosPosAlarmAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMosPosAlarmResponse `json:"alibaba_mos_pos_alarm_response,omitempty"`
}

type AlibabaMosPosAlarmResponse struct {

    // errMsg
    ErrorMessage   string `json:"error_message,omitempty"`

    // errCode
    SubErrCode   int64 `json:"sub_err_code,omitempty"`

    // success/false
    Result   bool `json:"result,omitempty"`

}
