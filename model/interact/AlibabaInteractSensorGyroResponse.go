package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
陀螺仪 APIResponse
alibaba.interact.sensor.gyro

客户端陀螺仪
*/
type AlibabaInteractSensorGyroAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorGyroResponse
}

type AlibabaInteractSensorGyroResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_gyro_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // return=0 表示正确
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
