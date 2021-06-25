package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
隐藏titleBar APIResponse
alibaba.interact.sensor.titlebarhide

隐藏titleBar
*/
type AlibabaInteractSensorTitlebarhideAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorTitlebarhideResponse `json:"alibaba_interact_sensor_titlebarhide_response,omitempty"`
}

type AlibabaInteractSensorTitlebarhideResponse struct {

    // return=0表示成功
    Result   string `json:"result,omitempty"`

}
