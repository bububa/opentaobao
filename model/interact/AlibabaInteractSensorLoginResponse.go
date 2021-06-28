package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆页面 APIResponse
alibaba.interact.sensor.login

获取登陆页面
*/
type AlibabaInteractSensorLoginAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_login_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0表示成功
    
    Result   string `json:"result,omitempty" xml:"