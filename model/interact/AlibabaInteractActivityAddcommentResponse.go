package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
微淘评论接口 APIResponse
alibaba.interact.activity.addcomment

发表评论，并返回楼层
*/
type AlibabaInteractActivityAddcommentAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_activity_addcomment_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"