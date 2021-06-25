package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆页面 APIResponse
alibaba.interact.sensor.login

获取登陆页面
*/
type AlibabaInteractSensorLoginAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorLoginResponse `json:"alibaba_interact_sensor_login_response,omitempty"`
}

type AlibabaInteractSensorLoginResponse struct {

    // return=0表示成功
    Result   string `json:"result,omitempty"`

}
