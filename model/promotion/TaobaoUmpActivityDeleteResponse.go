package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除营销活动 APIResponse
taobao.ump.activity.delete

删除营销活动。对应的活动详情等将会被全部删除。
*/
type TaobaoUmpActivityDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ump_activity_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"