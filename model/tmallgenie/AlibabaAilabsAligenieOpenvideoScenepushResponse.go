package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
视频单集场景接入API APIResponse
alibaba.ailabs.aligenie.openvideo.scenepush

视频单集场景接入API
*/
type AlibabaAilabsAligenieOpenvideoScenepushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieOpenvideoScenepushResponse
}

type AlibabaAilabsAligenieOpenvideoScenepushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_openvideo_scenepush_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 状态码
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 描述
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
