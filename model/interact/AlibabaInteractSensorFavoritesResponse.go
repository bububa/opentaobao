package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
手淘开放收藏夹鉴权接口 APIResponse
alibaba.interact.sensor.favorites

手淘开放鉴权专用接口，无数据输出输入，仅用于鉴权。
*/
type AlibabaInteractSensorFavoritesAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorFavoritesResponse `json:"alibaba_interact_sensor_favorites_response,omitempty"` 
    AlibabaInteractSensorFavoritesResponse
}

/* model for simplify = false
type AlibabaInteractSensorFavoritesResponse struct {

    // r=0
    
    R   string `json:"r,omitempty"`
    

}
*/

type AlibabaInteractSensorFavoritesResponse struct {

    // r=0
    R   string `json:"r,omitempty"`

}
