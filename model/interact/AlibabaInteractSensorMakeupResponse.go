package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
美妆虚拟试装 APIResponse
alibaba.interact.sensor.makeup

手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
*/
type AlibabaInteractSensorMakeupAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_makeup_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 非restAPI，为jsapi  result=0
    
    Result   string `json:"result,omitempty" xml:"