package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
美妆虚拟试装 APIResponse
alibaba.interact.sensor.makeup

手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
*/
type AlibabaInteractSensorMakeupAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorMakeupResponse `json:"alibaba_interact_sensor_makeup_response,omitempty"`
}

type AlibabaInteractSensorMakeupResponse struct {

    // 非restAPI，为jsapi  result=0
    Result   string `json:"result,omitempty"`

}
