package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
popwindow APIResponse
alibaba.interact.sensor.popwindow

popwindow
*/
type AlibabaInteractSensorPopwindowAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_popwindow_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0 表示成功
    
    Result   string `json:"result,omitempty" xml:"