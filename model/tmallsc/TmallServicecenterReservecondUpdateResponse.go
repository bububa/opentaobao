package tmallsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
主动预约条件更新 APIResponse
tmall.servicecenter.reservecond.update

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondUpdateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterReservecondUpdateResponse
}

type TmallServicecenterReservecondUpdateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_reservecond_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    MsgSuccess   bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`

    
    // 返回信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 返回code
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
