package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘开放收藏夹鉴权接口 APIResponse
alibaba.interact.sensor.favorites

手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
*/
type AlibabaInteractSensorFavoritesAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_favorites_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // r=0
    
    R   string `json:"r,omitempty" xml:"