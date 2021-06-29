package youkuott

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
dvb首次安装订单反馈 APIResponse
youku.ott.dvb.firstorder.feedback

dvb首次安装订单反馈
*/
type YoukuOttDvbFirstorderFeedbackAPIResponse struct {
    model.CommonResponse
    YoukuOttDvbFirstorderFeedbackResponse
}

type YoukuOttDvbFirstorderFeedbackResponse struct {
    XMLName xml.Name `xml:"youku_ott_dvb_firstorder_feedback_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功 true:成功 false:失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误消息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
