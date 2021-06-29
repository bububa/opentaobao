package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵内容接入标准接口 APIResponse
alibaba.ailabs.aligenie.opencontent.push

第三方内容接入天猫精灵内容库，供相关技能使用
*/
type AlibabaAilabsAligenieOpencontentPushAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieOpencontentPushResponse
}

type AlibabaAilabsAligenieOpencontentPushResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_opencontent_push_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaAilabsAligenieOpencontentPushResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
