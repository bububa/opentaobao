package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
pos故障报警 APIResponse
alibaba.mos.pos.alarm

故障报警
*/
type AlibabaMosPosAlarmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mos_pos_alarm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errMsg
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"