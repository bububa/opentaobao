package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频专辑场景接入接口 APIResponse
alibaba.ailabs.aligenie.openvideoalbum.scenepush

视频专辑场景接入接口
*/
type AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieOpenvideoalbumScenepushResponse
}

type AlibabaAilabsAligenieOpenvideoalbumScenepushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideoalbum_scenepush_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 状态码
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 描述
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
