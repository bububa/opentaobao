package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容库视频合辑数据推送接口 APIResponse
alibaba.ailabs.aligenie.videoalbum.push

三方内容合作厂商可将视频辑数据通过此接口推送至天猫精灵内容库接入中，供天猫精灵使用
*/
type AlibabaAilabsAligenieVideoalbumPushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieVideoalbumPushResponse
}

type AlibabaAilabsAligenieVideoalbumPushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_videoalbum_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 状态码
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 描述
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
