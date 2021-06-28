package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘拉起旺旺接口 APIResponse
alibaba.interact.sensor.wangwang

手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。
*/
type AlibabaInteractSensorWangwangAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_wangwang_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=0
    
    Result   string `json:"result,omitempty" xml:"