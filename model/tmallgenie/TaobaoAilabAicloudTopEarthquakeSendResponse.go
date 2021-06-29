package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
地震局发送地震消息 APIResponse
taobao.ailab.aicloud.top.earthquake.send

地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户
*/
type TaobaoAilabAicloudTopEarthquakeSendAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopEarthquakeSendResponse
}

type TaobaoAilabAicloudTopEarthquakeSendResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_earthquake_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应结果
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 扩展字段
    
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`

    
    // 响应错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 响应错误码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
