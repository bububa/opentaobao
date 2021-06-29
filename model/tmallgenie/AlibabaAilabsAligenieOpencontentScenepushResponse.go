package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
音频场景接入接口 APIResponse
alibaba.ailabs.aligenie.opencontent.scenepush

天猫精灵音频挂靠场景接入
*/
type AlibabaAilabsAligenieOpencontentScenepushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieOpencontentScenepushResponse
}

type AlibabaAilabsAligenieOpencontentScenepushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_opencontent_scenepush_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 返回信息
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`

    
}
