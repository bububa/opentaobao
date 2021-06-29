package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容库视频分集数据推送接口 APIResponse
alibaba.ailabs.aligenie.openvideo.push

天猫精灵内容库视频分集数据推送接口
*/
type AlibabaAilabsAligenieOpenvideoPushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieOpenvideoPushResponse
}

type AlibabaAilabsAligenieOpenvideoPushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideo_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 状态码
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 描述
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
