package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
客户端震动 API返回值 
alibaba.interact.sensor.vibrate

客户端震动
*/
type AlibabaInteractSensorVibrateAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorVibrateAPIResponseModel
}

// 客户端震动 成功返回结果
type AlibabaInteractSensorVibrateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_vibrate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result=0表示成功
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
