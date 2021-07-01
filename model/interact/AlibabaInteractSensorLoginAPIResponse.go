package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取登陆页面 API返回值 
alibaba.interact.sensor.login

获取登陆页面
*/
type AlibabaInteractSensorLoginAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorLoginAPIResponseModel
}

// 获取登陆页面 成功返回结果
type AlibabaInteractSensorLoginAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_login_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // return=0表示成功
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
