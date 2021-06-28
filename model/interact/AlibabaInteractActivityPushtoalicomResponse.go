package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
小铺isv推广流量活动到流量聚乐部 APIResponse
alibaba.interact.activity.pushtoalicom

涉及到流量包的小铺isv，将活动推送到流量聚乐部
*/
type AlibabaInteractActivityPushtoalicomAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractActivityPushtoalicomResponse `json:"alibaba_interact_activity_pushtoalicom_response,omitempty"` 
    AlibabaInteractActivityPushtoalicomResponse
}

/* model for simplify = false
type AlibabaInteractActivityPushtoalicomResponse struct {

    // 推送成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AlibabaInteractActivityPushtoalicomResponse struct {

    // 推送成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
