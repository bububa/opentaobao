package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动详情 APIResponse
taobao.ump.detail.update

更新活动详情
*/
type TaobaoUmpDetailUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ump_detail_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"