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
    AlibabaMosPosAlarmResponse
}

type AlibabaMosPosAlarmResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_pos_alarm_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errMsg
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`

    
    // errCode
    
    SubErrCode   int64 `json:"sub_err_code,omitempty" xml:"sub_err_code,omitempty"`

    
    // success/false
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
