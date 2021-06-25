package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 APIResponse
alibaba.interact.sensor.glue

视频播放
*/
type AlibabaInteractSensorGlueAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorGlueResponse `json:"alibaba_interact_sensor_glue_response,omitempty"`
}

type AlibabaInteractSensorGlueResponse struct {

    // result=0
    Result   string `json:"result,omitempty"`

}
