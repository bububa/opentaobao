package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ISV互动应用活动关闭服务 APIResponse
alibaba.interact.activity.unregister

卖家在ISV互动应用中设置的活动主动关闭的服务
*/
type AlibabaInteractActivityUnregisterAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractActivityUnregisterResponse `json:"alibaba_interact_activity_unregister_response,omitempty"`
}

type AlibabaInteractActivityUnregisterResponse struct {

    // 关闭活动结果
    UnregisterResult   *AllsparkResult `json:"unregister_result,omitempty"`

}
