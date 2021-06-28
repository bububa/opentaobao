package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
社交组件 APIResponse
alibaba.interact.sensor.social

赞，评论 ，关注 新增接口
*/
type AlibabaInteractSensorSocialAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_social_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=1
    
    Result   string `json:"result,omitempty" xml:"