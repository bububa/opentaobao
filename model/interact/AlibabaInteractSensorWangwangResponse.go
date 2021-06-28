package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
手淘拉起旺旺接口 APIResponse
alibaba.interact.sensor.wangwang

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。
*/
type AlibabaInteractSensorWangwangAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorWangwangResponse `json:"alibaba_interact_sensor_wangwang_response,omitempty"` 
    AlibabaInteractSensorWangwangResponse
}

/* model for simplify = false
type AlibabaInteractSensorWangwangResponse struct {

    // result=0
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorWangwangResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
