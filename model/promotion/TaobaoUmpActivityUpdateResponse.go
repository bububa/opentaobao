package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改活动信息 APIResponse
taobao.ump.activity.update

修改营销活动
*/
type TaobaoUmpActivityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoUmpActivityUpdateResponse
}

type TaobaoUmpActivityUpdateResponse struct {
    XMLName xml.Name `xml:"ump_activity_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
