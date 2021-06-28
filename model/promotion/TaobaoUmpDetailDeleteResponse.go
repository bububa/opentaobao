package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除活动详情 APIResponse
taobao.ump.detail.delete

删除活动详情
*/
type TaobaoUmpDetailDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoUmpDetailDeleteResponse
}

type TaobaoUmpDetailDeleteResponse struct {
    XMLName xml.Name `xml:"ump_detail_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
