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
    AlibabaInteractSensorLoginResponse
}

type AlibabaInteractSensorLoginResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_login_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0表示成功
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
