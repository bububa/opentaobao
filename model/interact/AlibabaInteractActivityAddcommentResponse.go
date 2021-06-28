package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
微淘评论接口 APIResponse
alibaba.interact.activity.addcomment

发表评论，并返回楼层
*/
type AlibabaInteractActivityAddcommentAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractActivityAddcommentResponse `json:"alibaba_interact_activity_addcomment_response,omitempty"` 
    AlibabaInteractActivityAddcommentResponse
}

/* model for simplify = false
type AlibabaInteractActivityAddcommentResponse struct {

    // 返回成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 评论的楼层数
    
    Floor   int64 `json:"floor,omitempty"`
    

}
*/

type AlibabaInteractActivityAddcommentResponse struct {

    // 返回成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 评论的楼层数
    Floor   int64 `json:"floor,omitempty"`

}
