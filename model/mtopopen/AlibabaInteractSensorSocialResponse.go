package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
社交组件 APIResponse
alibaba.interact.sensor.social

赞，评论 ，关注 新增接口
*/
type AlibabaInteractSensorSocialAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractSensorSocialResponse `json:"alibaba_interact_sensor_social_response,omitempty"` 
    AlibabaInteractSensorSocialResponse
}

/* model for simplify = false
type AlibabaInteractSensorSocialResponse struct {

    // result=1
    
    Result   string `json:"result,omitempty"`
    

}
*/

type AlibabaInteractSensorSocialResponse struct {

    // result=1
    Result   string `json:"result,omitempty"`

}
