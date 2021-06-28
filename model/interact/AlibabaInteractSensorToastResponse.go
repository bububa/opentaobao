package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
toast APIResponse
alibaba.interact.sensor.toast

toast提示
*/
type AlibabaInteractSensorToastAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_toast_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0 表示成功
    
    Result   string `json:"result,omitempty" xml:"