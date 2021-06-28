package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
陀螺仪 APIResponse
alibaba.interact.sensor.gyro

客户端陀螺仪
*/
type AlibabaInteractSensorGyroAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorGyroResponse `json:"alibaba_interact_sensor_gyro_response,omitempty"` 
    AlibabaInteractSensorGyroResponse
}

/* model for simplify = false
type AlibabaInteractSensorGyroResponse struct {

    // return=0 表示正确
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorGyroResponse struct {

    // return=0 表示正确
    Result   string `json:"result,omitempty"`

}
